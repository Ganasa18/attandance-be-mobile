package service

import (
	"context"
	"ganasa18/attandance-be-mobile/internal/model/web"
)

type AuthService interface {
	Register(ctx context.Context, request web.UserCreateRequest) web.UserCreateRequest
	Login(ctx context.Context, request web.AuthLoginRequest) (web.AuthLoginResponseRequest, error)
	LoginWithOutOtp(ctx context.Context, request web.AuthLoginRequest) (web.AuthLoginResponseRequest, error)
}
