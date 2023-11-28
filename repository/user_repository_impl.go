package repository

import (
	"context"
	"database/sql"
	"fmt"
	"ganasa18/attandance-be-mobile/helper"
	"ganasa18/attandance-be-mobile/model/domain"
	"net/url"
	"strconv"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.UserModel, int) {

	defaultPage := 1
	defaultLimit := 10
	page := defaultPage
	limit := defaultLimit

	if queryParams, ok := ctx.Value("queryParams").(url.Values); ok {

		pageValue := queryParams.Get("page")
		limitValue := queryParams.Get("limit")

		if pageValue != "" {
			page, _ = strconv.Atoi(pageValue)
		}

		if limitValue != "" {
			limit, _ = strconv.Atoi(limitValue)
		}

	} else {
		fmt.Println("Failed to read query parameters from context")
	}
	offset := (page - 1) * limit

	// Query to get the total count of rows
	countQuery := "SELECT COUNT(*) FROM ms_users"
	var rowCount int
	err := tx.QueryRowContext(ctx, countQuery).Scan(&rowCount)
	helper.PanicIfError(err)

	SQL := "SELECT id, username, email, is_active, created_at, updated_at, deleted_at FROM ms_users LIMIT $1 OFFSET $2"
	rows, err := tx.QueryContext(ctx, SQL, limit, offset)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []domain.UserModel
	for rows.Next() {
		user := domain.UserModel{}
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.IsActive, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		helper.PanicIfError(err)
		users = append(users, user)
	}
	return users, rowCount
}
