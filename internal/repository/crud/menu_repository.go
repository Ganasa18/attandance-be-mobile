package repository

import (
	"context"
	"database/sql"
	"ganasa18/attandance-be-mobile/internal/model/domain"
)

type MenuRepository interface {
	CreateMenu(ctx context.Context, tx *sql.Tx, menu domain.MenuModel) domain.MenuModel
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.MenuModel, int)
}
