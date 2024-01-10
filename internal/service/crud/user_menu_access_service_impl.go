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

type UserMenuAccessServiceImpl struct {
	UserMenuAccessRepository repository.UserMenuAccessRepository
	DB                       *sql.DB
	Validate                 *validator.Validate
}

func NewUserMenuAccessService(userMenuAccessRepository repository.UserMenuAccessRepository, DB *sql.DB, validate *validator.Validate) UserMenuAccessService {
	return &UserMenuAccessServiceImpl{
		UserMenuAccessRepository: userMenuAccessRepository,
		DB:                       DB,
		Validate:                 validate,
	}
}

func (service *UserMenuAccessServiceImpl) CreateUserMenuAccess(ctx context.Context, request web.UserMenuCreateAccessResponseRequest) web.UserMenuAccessResponseRequest {
	// VALIDATE
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// TRANSACTION BEGIN
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	// CHECK TRANSACTION
	defer helper.CommitOrRollback(tx)

	userMenuAccess := domain.UserMenuAccessModel{
		MenuId:       request.MenuId,
		UserId:       request.UserId,
		RoleAccessId: request.RoleAccessId,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	userMenuAccess = service.UserMenuAccessRepository.CreateUserMenuAccess(ctx, tx, userMenuAccess)

	return modelhelper.ToUserMenuAccessResponse(userMenuAccess)
}

func (service *UserMenuAccessServiceImpl) GetUserMenuAccess(ctx context.Context) []web.UserMenuAccessByUserIdResponseRequest {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	// CHECK TRANSACTION
	defer helper.CommitOrRollback(tx)
	// LOGIC
	userMenus := service.UserMenuAccessRepository.GetUserMenuAccess(ctx, tx)
	return modelhelper.ToUserMenusJoinResponses(userMenus)
}
