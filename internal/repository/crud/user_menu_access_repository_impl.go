package repository

import (
	"context"
	"database/sql"
	"ganasa18/attandance-be-mobile/internal/helper"
	"ganasa18/attandance-be-mobile/internal/model/domain"
	"log"
)

type UserMenuAccessRepositoryImpl struct{}

func NewUserMenuAccessRepository() UserMenuAccessRepository {
	return &UserMenuAccessRepositoryImpl{}
}

func (repository *UserMenuAccessRepositoryImpl) CreateUserMenuAccess(ctx context.Context, tx *sql.Tx, userMenuAccess domain.UserMenuAccessModel) domain.UserMenuAccessModel {

	SQL := "INSERT INTO ms_user_menu_access (menu_id, user_id, role_access_id, created_at, updated_at) values($1 , $2, $3, $4, $5) RETURNING id"
	var lastInsertID int
	values := []interface{}{userMenuAccess.MenuId, userMenuAccess.UserId, userMenuAccess.RoleAccessId, userMenuAccess.CreatedAt, userMenuAccess.UpdatedAt}
	err := tx.QueryRowContext(ctx, SQL, values...).Scan(&lastInsertID)
	helper.PanicIfError(err)
	// LOGGING
	log.Printf("Executing SQL: %s with values: %v\n", SQL, helper.FormatArraySplitByComma(values))
	userMenuAccess.Id = lastInsertID
	return userMenuAccess
}

func (repository *UserMenuAccessRepositoryImpl) GetUserMenuAccess(ctx context.Context, tx *sql.Tx) []domain.UserMenuAccessJoin {

	SQL := `
    SELECT 
        ms_user_menu_access.id AS uma_id, 
        ms_user_menu_access.menu_id, 
        ms_user_menu_access.user_id, 
        ms_user_menu_access.role_access_id, 
        ms_role.id AS ms_role_id, 
        ms_role.role_name, 
        ms_users.id AS msu_id, 
        ms_menu.id AS msm_id, 
        ms_menu.menu_name, 
        ms_menu.title_menu, 
        ms_menu.path,
		ms_menu.is_submenu,
		ms_menu.parent_name,
        ms_role_access.role_id
    FROM 
        ms_user_menu_access
    INNER JOIN 
        ms_users ON ms_users.id = ms_user_menu_access.user_id
    INNER JOIN 
        ms_role_access ON ms_role_access.role_id = ms_user_menu_access.role_access_id
    INNER JOIN 
        ms_menu ON ms_menu.id = ms_user_menu_access.menu_id
    INNER JOIN 
        ms_role ON ms_role.id = ms_role_access.role_id
`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	log.Printf("Executing SQL: %s", SQL)

	defer rows.Close()
	var userAccessMenu []domain.UserMenuAccessJoin
	for rows.Next() {
		userMenu := domain.UserMenuAccessJoin{}

		// var umaID, menuID, userID, roleAccessID, msRoleID, msuID, msmID, roleID int
		// var roleName, menuName, titleMenu, path string

		// err := rows.Scan(
		// 	&umaID, &menuID, &userID, &roleAccessID, &msRoleID, &roleName, &msuID, &msmID,
		// 	&menuName, &titleMenu, &path, &roleID,
		// )

		// log.Printf("Row: uma_id=%d, menu_id=%d, user_id=%d, role_access_id=%d, ms_role_id=%d, role_name=%s, msu_id=%d, msm_id=%d, menu_name=%s, title_menu=%s, path=%s, role_id=%d",
		// 	umaID, menuID, userID, roleAccessID, msRoleID, roleName, msuID, msmID, menuName, titleMenu, path, roleID,
		// )

		err := rows.Scan(
			&userMenu.UmaID,
			&userMenu.MenuID,
			&userMenu.UserID,
			&userMenu.RoleAccessID,
			&userMenu.MsRoleID,
			&userMenu.RoleName,
			&userMenu.MsUserID,
			&userMenu.MsMenuID,
			&userMenu.MenuName,
			&userMenu.TitleMenu,
			&userMenu.Path,
			&userMenu.IsSubmenu,
			&userMenu.ParentName,
			&userMenu.RoleID,
		)

		// Print the scanned values
		// log.Printf("UmaID: %d, MenuID: %d, UserID: %d, RoleAccessID: %d, MsRoleID: %d, MsUserID: %d, MsMenuID: %d, RoleID: %d, RoleName: %s, MenuName: %s, TitleMenu: %s, Path: %s",
		// 	userMenu.UmaID, userMenu.MenuID, userMenu.UserID, userMenu.RoleAccessID, userMenu.MsRoleID, userMenu.MsUserID, userMenu.MsMenuID, userMenu.RoleID, userMenu.RoleName, userMenu.MenuName, userMenu.TitleMenu, userMenu.Path)

		helper.PanicIfError(err)
		userAccessMenu = append(userAccessMenu, userMenu)
	}

	return userAccessMenu

}
