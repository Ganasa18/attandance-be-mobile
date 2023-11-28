package modelhelper

import (
	"ganasa18/attandance-be-mobile/model/domain"
	"ganasa18/attandance-be-mobile/model/web"
)

func ToUserResponse(user domain.UserModel) web.UserResponseRequest {
	return web.UserResponseRequest{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}

func ToUserResponses(users []domain.UserModel, rowCount int) ([]web.UserResponseRequest, int) {
	var userResponse []web.UserResponseRequest
	for _, user := range users {
		userResponse = append(userResponse, ToUserResponse(user))
	}
	return userResponse, rowCount
}
