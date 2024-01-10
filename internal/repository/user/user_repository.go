package repository

import (
	"context"
	"database/sql"
	"ganasa18/attandance-be-mobile/internal/model/domain"
)

type UserRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.UserModel, int)
	FindOne(ctx context.Context, tx *sql.Tx, userId int) (domain.UserModel, error)
}
