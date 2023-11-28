package repository

import (
	"context"
	"database/sql"
	"ganasa18/attandance-be-mobile/model/domain"
)

type UserRepository interface {
	// Register(ctx context.Context, tx *sql.Tx, user domain.UserModel) domain.UserModel
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.UserModel, int)
}
