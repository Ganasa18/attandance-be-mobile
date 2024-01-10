package service

import (
	"context"
	"database/sql"
	"ganasa18/attandance-be-mobile/internal/helper"
	"ganasa18/attandance-be-mobile/internal/helper/modelhelper"
	"ganasa18/attandance-be-mobile/internal/model/domain"
	"ganasa18/attandance-be-mobile/internal/model/web"
	repository "ganasa18/attandance-be-mobile/internal/repository/auth"

	"time"

	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewAuthService(authRepository repository.AuthRepository, DB *sql.DB, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		AuthRepository: authRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *AuthServiceImpl) Register(ctx context.Context, request web.UserCreateRequest) web.UserCreateRequest {
	// VALIDATE
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// TRANSACTION BEGIN
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	// CHECK TRANSACTION
	defer helper.CommitOrRollback(tx)

	password := []byte(request.Password)
	// Hashing the password with the default cost of 10
	hashedPassword, errHashedPassword := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	helper.PanicIfError(errHashedPassword)

	// GENERATE UUID
	uniqueID := helper.GenerateUUID()

	// LOGIC
	register := domain.UserRegisterModel{
		Username:     request.Username,
		Email:        request.Email,
		UserUniqueId: uniqueID,
		Password:     string(hashedPassword),
		UserRole:     request.UserRole,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	register = service.AuthRepository.Register(ctx, tx, register)
	return modelhelper.ToUserRegisterResponse(register)
}

func (service *AuthServiceImpl) Login(ctx context.Context, request web.AuthLoginRequest) (web.AuthLoginResponseRequest, error) {
	// VALIDATE
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	// TRANSACTION BEGIN
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	// CHECK TRANSACTION
	defer helper.CommitOrRollback(tx)
	// LOGIC
	login := domain.UserModel{
		Email:    request.Email,
		Password: request.Password,
	}
	userData, err := service.AuthRepository.Login(ctx, tx, login)
	return modelhelper.ToAuthLoginResponse(userData, err)

}

func (service *AuthServiceImpl) LoginWithOutOtp(ctx context.Context, request web.AuthLoginRequest) (web.AuthLoginResponseRequest, error) {
	// VALIDATE
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	// TRANSACTION BEGIN
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	// CHECK TRANSACTION
	defer helper.CommitOrRollback(tx)
	// LOGIC
	login := domain.UserModel{
		Email:    request.Email,
		Password: request.Password,
	}
	userData, err := service.AuthRepository.Login(ctx, tx, login)
	return modelhelper.ToAuthLoginResponse(userData, err)

}
