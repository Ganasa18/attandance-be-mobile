package repository

import (
	"context"
	"database/sql"
	"ganasa18/attandance-be-mobile/internal/model/domain"
)

type RoleRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.RoleModel, int)
	CreateRole(ctx context.Context, tx *sql.Tx, role domain.RoleModel) domain.RoleModel
}
