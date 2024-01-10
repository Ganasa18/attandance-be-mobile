package service

import (
	"context"
	"ganasa18/attandance-be-mobile/internal/model/web"
)

type UserMenuAccessService interface {
	CreateUserMenuAccess(ctx context.Context, request web.UserMenuCreateAccessResponseRequest) web.UserMenuAccessResponseRequest
	GetUserMenuAccess(ctx context.Context) []web.UserMenuAccessByUserIdResponseRequest
}
