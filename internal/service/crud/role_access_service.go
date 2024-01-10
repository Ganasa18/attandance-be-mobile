package service

import (
	"context"
	"ganasa18/attandance-be-mobile/internal/model/web"
)

type RoleAccessService interface {
	FindAll(ctx context.Context) ([]web.RoleAccessMasterResponseRequest, int)
	CreateRoleAccess(ctx context.Context, request web.RoleAccessMasterCreateResponseRequest) web.RoleAccessMasterResponseRequest
}
