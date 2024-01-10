package service

import (
	"context"
	"database/sql"
	"ganasa18/attandance-be-mobile/internal/helper"
	"ganasa18/attandance-be-mobile/internal/helper/modelhelper"
	"ganasa18/attandance-be-mobile/internal/model/domain"
	"ganasa18/attandance-be-mobile/internal/model/web"
	repository "ganasa18/attandance-be-mobile/internal/repository/crud"
	"time"

	"github.com/go-playground/validator"
)

type RoleAccessServiceImpl struct {
	RoleAccessRepository repository.RoleAccessRepository
	DB                   *sql.DB
	Validate             *validator.Validate
}

func NewRoleAccessService(roleAccessRepository repository.RoleAccessRepository, DB *sql.DB, validate *validator.Validate) RoleAccessService {
	return &RoleAccessServiceImpl{
		RoleAccessRepository: roleAccessRepository,
		DB:                   DB,
		Validate:             validate,
	}
}

func (service *RoleAccessServiceImpl) FindAll(ctx context.Context) ([]web.RoleAccessMasterResponseRequest, int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	// CHECK TRANSACTION
	defer helper.CommitOrRollback(tx)
	// LOGIC
	roleAccess, rowCount := service.RoleAccessRepository.FindAll(ctx, tx)

	return modelhelper.ToRoleAccessResponses(roleAccess, rowCount)
}

func (service *RoleAccessServiceImpl) CreateRoleAccess(ctx context.Context, request web.RoleAccessMasterCreateResponseRequest) web.RoleAccessMasterResponseRequest {
	// VALIDATE
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	// TRANSACTION BEGIN
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	// CHECK TRANSACTION
	defer helper.CommitOrRollback(tx)
	roleAccess := domain.RoleAccessModel{
		RoleId:    request.RoleId,
		Create:    request.Create,
		Read:      request.Read,
		Update:    request.Update,
		Delete:    request.Delete,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	roleAccess = service.RoleAccessRepository.CreateRoleAccess(ctx, tx, roleAccess)
	return modelhelper.ToRoleAccessResponse(roleAccess)
}
