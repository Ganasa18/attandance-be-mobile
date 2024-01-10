package service

import (
	"context"
	"ganasa18/attandance-be-mobile/internal/model/web"
)

type RoleService interface {
	CreateRole(ctx context.Context, request web.RoleMasterCreateResponseRequest) web.RoleMasterResponseRequest
	FindAll(ctx context.Context) ([]web.RoleMasterResponseRequest, int)
}
