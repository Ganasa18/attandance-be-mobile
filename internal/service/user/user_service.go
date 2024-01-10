package service

import (
	"context"
	"ganasa18/attandance-be-mobile/internal/model/web"
)

type UserService interface {
	FindAll(ctx context.Context) ([]web.UserResponseRequest, int)
	FindOne(ctx context.Context, userId int) web.UserResponseRequest
}
