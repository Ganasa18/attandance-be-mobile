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

type MenuServiceImpl struct {
	MenuRepository repository.MenuRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewMenuService(menuRepository repository.MenuRepository, DB *sql.DB, validate *validator.Validate) MenuService {
	return &MenuServiceImpl{
		MenuRepository: menuRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *MenuServiceImpl) CreateMenu(ctx context.Context, request web.MenuMasterCreateResponseRequest) web.MenuMasterResponseRequest {
	// VALIDATE
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// TRANSACTION BEGIN
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	// CHECK TRANSACTION
	defer helper.CommitOrRollback(tx)

	menu := domain.MenuModel{
		MenuName:   request.MenuName,
		TitleMenu:  request.TitleMenu,
		Path:       request.Path,
		IsSubmenu:  request.IsSubmenu,
		ParentName: request.ParentName,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	menu = service.MenuRepository.CreateMenu(ctx, tx, menu)
	return modelhelper.ToMenuResponse(menu)
}

func (service *MenuServiceImpl) FindAll(ctx context.Context) ([]web.MenuMasterResponseRequest, int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	// CHECK TRANSACTION
	defer helper.CommitOrRollback(tx)
	// LOGIC
	menu, rowCount := service.MenuRepository.FindAll(ctx, tx)
	return modelhelper.ToMenuResponses(menu, rowCount)
}
