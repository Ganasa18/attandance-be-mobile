package service

import (
	"context"
	"ganasa18/attandance-be-mobile/model/web"
)

type UserService interface {
	// Register(cctx context.Context, request web.UserCreateRequest) web.UserResponseRequest
	FindAll(ctx context.Context) ([]web.UserResponseRequest, int)
}
