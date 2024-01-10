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

type RoleServiceImpl struct {
	RoleRepository repository.RoleRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewRoleService(roleRepository repository.RoleRepository, DB *sql.DB, validate *validator.Validate) RoleService {
	return &RoleServiceImpl{
		RoleRepository: roleRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *RoleServiceImpl) CreateRole(ctx context.Context, request web.RoleMasterCreateResponseRequest) web.RoleMasterResponseRequest {
	// VALIDATE
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// TRANSACTION BEGIN
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	// CHECK TRANSACTION
	defer helper.CommitOrRollback(tx)

	role := domain.RoleModel{
		Rolename:  request.Rolename,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	role = service.RoleRepository.CreateRole(ctx, tx, role)
	return modelhelper.ToRoleResponse(role)
}

func (service *RoleServiceImpl) FindAll(ctx context.Context) ([]web.RoleMasterResponseRequest, int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	// CHECK TRANSACTION
	defer helper.CommitOrRollback(tx)
	// LOGIC
	roles, rowCount := service.RoleRepository.FindAll(ctx, tx)
	return modelhelper.ToRoleResponses(roles, rowCount)
}
