package service

import (
	"context"
	"database/sql"
	"ganasa18/attandance-be-mobile/internal/exception"
	"ganasa18/attandance-be-mobile/internal/helper"
	"ganasa18/attandance-be-mobile/internal/helper/modelhelper"
	"ganasa18/attandance-be-mobile/internal/model/web"
	repository "ganasa18/attandance-be-mobile/internal/repository/user"

	"github.com/go-playground/validator"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) FindAll(ctx context.Context) ([]web.UserResponseRequest, int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	// CHECK TRANSACTION
	defer helper.CommitOrRollback(tx)
	// LOGIC
	users, rowCount := service.UserRepository.FindAll(ctx, tx)
	return modelhelper.ToUserResponses(users, rowCount)
}

func (service *UserServiceImpl) FindOne(ctx context.Context, userId int) web.UserResponseRequest {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	// CHECK TRANSACTION
	defer helper.CommitOrRollback(tx)

	// LOGIC
	user, err := service.UserRepository.FindOne(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return modelhelper.ToUserResponse(user)
}
