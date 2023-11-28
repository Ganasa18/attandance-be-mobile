package service

import (
	"context"
	"database/sql"
	"ganasa18/attandance-be-mobile/helper"
	"ganasa18/attandance-be-mobile/helper/modelhelper"
	"ganasa18/attandance-be-mobile/model/web"
	"ganasa18/attandance-be-mobile/repository"

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
