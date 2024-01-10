package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"ganasa18/attandance-be-mobile/internal/helper"
	"ganasa18/attandance-be-mobile/internal/model/domain"
	"log"
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

	SQL := `
			SELECT
				ms_users.id, user_unique_id, username, 
				email, is_active, user_role, 
				ms_users.created_at, ms_users.updated_at, ms_users.deleted_at,
				ms_role.id AS role_id, ms_role.role_name  
			FROM 
				ms_users
			LEFT JOIN
				ms_role ON ms_role.id = ms_users.user_role
			LIMIT $1 OFFSET $2`
	values := []interface{}{limit, offset}
	rows, err := tx.QueryContext(ctx, SQL, values...)
	helper.PanicIfError(err)

	// LOGGING
	log.Printf("Executing SQL: %s with values: %v\n", SQL, helper.FormatArraySplitByComma(values))

	defer rows.Close()

	var users []domain.UserModel
	for rows.Next() {
		user := domain.UserModel{}
		err := rows.Scan(&user.Id, &user.UserUniqueId, &user.Username, &user.Email, &user.IsActive, &user.UserRole, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt, &user.RoleId, &user.RoleName)

		helper.PanicIfError(err)
		users = append(users, user)
	}
	return users, rowCount
}

func (repository *UserRepositoryImpl) FindOne(ctx context.Context, tx *sql.Tx, userId int) (domain.UserModel, error) {

	SQL := "SELECT id, username, email, is_active, created_at, updated_at FROM ms_users WHERE id = $1"
	values := []interface{}{userId}
	rows, err := tx.QueryContext(ctx, SQL, values...)
	helper.PanicIfError(err)
	defer rows.Close()

	// LOGGING
	log.Printf("Executing SQL: %s with values: %v\n", SQL, helper.FormatArraySplitByComma(values))

	user := domain.UserModel{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}
