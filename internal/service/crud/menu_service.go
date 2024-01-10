package service

import (
	"context"
	"ganasa18/attandance-be-mobile/internal/model/web"
)

type MenuService interface {
	CreateMenu(ctx context.Context, request web.MenuMasterCreateResponseRequest) web.MenuMasterResponseRequest
	FindAll(ctx context.Context) ([]web.MenuMasterResponseRequest, int)
}
