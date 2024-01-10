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

type MenuRepositoryImpl struct{}

func NewMenuRepository() MenuRepository {
	return &MenuRepositoryImpl{}
}

func (repository *MenuRepositoryImpl) CreateMenu(ctx context.Context, tx *sql.Tx, menu domain.MenuModel) domain.MenuModel {

	SQL := "INSERT INTO ms_menu (menu_name, title_menu, path, is_submenu, parent_name, created_at, updated_at) values($1, $2, $3, $4, $5, $6 ,$7 ) RETURNING id"
	var lastInsertID int
	values := []interface{}{menu.MenuName, menu.TitleMenu, menu.Path, menu.IsSubmenu, menu.ParentName, menu.CreatedAt, menu.UpdatedAt}
	err := tx.QueryRowContext(ctx, SQL, values...).Scan(&lastInsertID)
	helper.PanicIfError(err)
	menu.Id = lastInsertID
	// LOGGING
	log.Printf("Executing SQL: %s with values: %v\n", SQL, helper.FormatArraySplitByComma(values))
	return menu
}

func (repository *MenuRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.MenuModel, int) {
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
	countQuery := "SELECT COUNT(*) FROM ms_menu"
	var rowCount int
	err := tx.QueryRowContext(ctx, countQuery).Scan(&rowCount)
	helper.PanicIfError(err)

	SQL := "SELECT id, menu_name, title_menu, path, is_submenu, parent_name, created_at, updated_at FROM ms_menu LIMIT $1 OFFSET $2"
	values := []interface{}{limit, offset}
	rows, err := tx.QueryContext(ctx, SQL, values...)
	helper.PanicIfError(err)

	// LOGGING
	log.Printf("Executing SQL: %s with values: %v\n", SQL, helper.FormatArraySplitByComma(values))

	defer rows.Close()

	var menus []domain.MenuModel
	for rows.Next() {
		menu := domain.MenuModel{}
		err := rows.Scan(&menu.Id, &menu.MenuName, &menu.TitleMenu, &menu.Path, &menu.IsSubmenu, &menu.ParentName, &menu.CreatedAt, &menu.UpdatedAt)
		helper.PanicIfError(err)
		menus = append(menus, menu)
	}
	return menus, rowCount
}
