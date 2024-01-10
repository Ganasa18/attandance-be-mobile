package repository

import (
	"context"
	"database/sql"
	"ganasa18/attandance-be-mobile/internal/model/domain"
)

type RoleAccessRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.RoleAccessModel, int)
	CreateRoleAccess(ctx context.Context, tx *sql.Tx, roleAccess domain.RoleAccessModel) domain.RoleAccessModel
}
