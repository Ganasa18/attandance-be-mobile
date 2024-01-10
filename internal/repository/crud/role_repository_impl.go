package repository

import (
	"context"
	"database/sql"
	"fmt"
	"ganasa18/attandance-be-mobile/internal/helper"
	"ganasa18/attandance-be-mobile/internal/model/domain"
	"log"
	"net/url"
	"strconv"
)

type RoleRepositoryImpl struct{}

func NewRoleRepository() RoleRepository {
	return &RoleRepositoryImpl{}
}

func (repository *RoleRepositoryImpl) CreateRole(ctx context.Context, tx *sql.Tx, role domain.RoleModel) domain.RoleModel {

	SQL := "INSERT INTO ms_role (role_name, created_at, updated_at) values($1, $2, $3) RETURNING id"
	var lastInsertID int
	values := []interface{}{role.Rolename, role.CreatedAt, role.UpdatedAt}
	err := tx.QueryRowContext(ctx, SQL, values...).Scan(&lastInsertID)
	helper.PanicIfError(err)
	role.Id = lastInsertID
	// LOGGING
	log.Printf("Executing SQL: %s with values: %v\n", SQL, helper.FormatArraySplitByComma(values))
	return role
}

func (repository *RoleRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.RoleModel, int) {
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
	countQuery := "SELECT COUNT(*) FROM ms_role"
	var rowCount int
	err := tx.QueryRowContext(ctx, countQuery).Scan(&rowCount)
	helper.PanicIfError(err)

	SQL := "SELECT id, role_name, created_at, updated_at FROM ms_role LIMIT $1 OFFSET $2"
	values := []interface{}{limit, offset}
	rows, err := tx.QueryContext(ctx, SQL, values...)
	helper.PanicIfError(err)

	// LOGGING
	log.Printf("Executing SQL: %s with values: %v\n", SQL, helper.FormatArraySplitByComma(values))

	defer rows.Close()

	var roles []domain.RoleModel
	for rows.Next() {
		role := domain.RoleModel{}
		err := rows.Scan(&role.Id, &role.Rolename, &role.CreatedAt, &role.UpdatedAt)
		helper.PanicIfError(err)
		roles = append(roles, role)
	}
	return roles, rowCount
}
