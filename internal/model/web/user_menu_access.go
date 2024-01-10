package web

import "time"

type UserMenuAccessResponseRequest struct {
	Id           int       `json:"id"`
	MenuId       int       ` json:"menu_id"`
	UserId       int       `json:"user_id"`
	RoleAccessId int       `json:"role_access_id"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type UserMenuCreateAccessResponseRequest struct {
	Id           int       `json:"id"`
	MenuId       int       `validate:"required" json:"menu_id"`
	UserId       int       `validate:"required" json:"user_id"`
	RoleAccessId int       `validate:"required" json:"role_access_id"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type UserMenuAccessByUserIdResponseRequest struct {
	// Id        int `json:"id"`
	UmaID        int `json:"id"`
	MenuID       int `json:"menu_id"`
	UserID       int `json:"user_id"`
	RoleAccessID int `json:"role_access_id"`
	// MsRoleID     int    `json:"ms_role_id"`
	// MsUserID     int    `json:"ms_user_id"`
	// MsMenuID     int    `json:"ms_menu_id"`
	RoleID     int    `json:"role_id"`
	RoleName   string `json:"role_name"`
	MenuName   string `json:"menu_name"`
	TitleMenu  string `json:"title_menu"`
	Path       string `json:"path"`
	IsSubmenu  bool   `json:"is_submenu"`
	ParentName string `json:"parent_name"`
}
