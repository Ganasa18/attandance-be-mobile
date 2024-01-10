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

type RoleAccessRepositoryImpl struct{}

func NewRoleAccessRepository() RoleAccessRepository {
	return &RoleAccessRepositoryImpl{}
}

func (repository *RoleAccessRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.RoleAccessModel, int) {
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
	countQuery := "SELECT COUNT(*) FROM ms_role_access"
	var rowCount int
	err := tx.QueryRowContext(ctx, countQuery).Scan(&rowCount)
	helper.PanicIfError(err)

	SQL := `SELECT id, role_id, mra.create, read, update, delete, created_at, updated_at FROM ms_role_access mra LIMIT $1 OFFSET $2`
	values := []interface{}{limit, offset}
	rows, err := tx.QueryContext(ctx, SQL, limit, offset)
	helper.PanicIfError(err)

	// LOGGING
	log.Printf("Executing SQL: %s with values: %v\n", SQL, helper.FormatArraySplitByComma(values))

	defer rows.Close()
	var roleAccesss []domain.RoleAccessModel
	for rows.Next() {
		roleAccess := domain.RoleAccessModel{}
		err := rows.Scan(
			&roleAccess.Id,
			&roleAccess.RoleId,
			&roleAccess.Create,
			&roleAccess.Read,
			&roleAccess.Update,
			&roleAccess.Delete,
			&roleAccess.CreatedAt,
			&roleAccess.UpdatedAt,
		)

		helper.PanicIfError(err)
		roleAccesss = append(roleAccesss, roleAccess)
	}

	return roleAccesss, rowCount
}

func (repository *RoleAccessRepositoryImpl) CreateRoleAccess(ctx context.Context, tx *sql.Tx, roleAccess domain.RoleAccessModel) domain.RoleAccessModel {

	SQL := "INSERT INTO ms_role_access (role_id, \"create\", \"read\", \"update\", \"delete\", created_at, updated_at) values($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	var lastInsertID int
	values := []interface{}{roleAccess.RoleId, roleAccess.Create, roleAccess.Read, roleAccess.Update, roleAccess.Delete, roleAccess.CreatedAt, roleAccess.UpdatedAt}
	err := tx.QueryRowContext(ctx, SQL, values...).Scan(&lastInsertID)
	helper.PanicIfError(err)
	roleAccess.Id = lastInsertID
	// LOGGING
	log.Printf("Executing SQL: %s with values: %v\n", SQL, helper.FormatArraySplitByComma(values))
	return roleAccess
}
