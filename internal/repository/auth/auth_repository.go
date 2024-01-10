package repository

import (
	"context"
	"database/sql"
	"ganasa18/attandance-be-mobile/internal/model/domain"
)

type AuthRepository interface {
	Register(ctx context.Context, tx *sql.Tx, user domain.UserRegisterModel) domain.UserRegisterModel
	Login(ctx context.Context, tx *sql.Tx, user domain.UserModel) (domain.UserModel, error)
	LoginWithOutOtp(ctx context.Context, tx *sql.Tx, user domain.UserModel) (domain.UserModel, error)
}
