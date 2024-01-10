package modelhelper

import (
	"ganasa18/attandance-be-mobile/internal/model/domain"
	"ganasa18/attandance-be-mobile/internal/model/web"
)

func ToAuthLoginResponse(user domain.UserModel, errorData error) (web.AuthLoginResponseRequest, error) {

	loginResponse := web.AuthLoginResponseRequest{
		Id:           user.Id,
		UserUniqueId: user.UserUniqueId,
		Username:     user.Username,
		Email:        user.Email,
		IsActive:     user.IsActive,
		Token:        user.Token,
		UserRole:     user.UserRole,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}

	return loginResponse, errorData
}

func ToUserRegisterResponse(user domain.UserRegisterModel) web.UserCreateRequest {
	return web.UserCreateRequest{
		Id:           user.Id,
		UserUniqueId: user.UserUniqueId,
		Username:     user.Username,
		Email:        user.Email,
		Password:     user.Password,
		UserRole:     user.UserRole,
		IsActive:     user.IsActive,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}
