package repository

import (
	"context"
	"database/sql"
	"ganasa18/attandance-be-mobile/internal/model/domain"
)

type UserMenuAccessRepository interface {
	CreateUserMenuAccess(ctx context.Context, tx *sql.Tx, userMenuAccess domain.UserMenuAccessModel) domain.UserMenuAccessModel
	GetUserMenuAccess(ctx context.Context, tx *sql.Tx) []domain.UserMenuAccessJoin
}
