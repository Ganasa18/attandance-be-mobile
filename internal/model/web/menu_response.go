package web

import "time"

type MenuMasterResponseRequest struct {
	Id         int       `json:"id"`
	MenuName   string    `json:"menu_name"`
	TitleMenu  string    `json:"title_menu"`
	Path       string    `json:"path"`
	IsSubmenu  bool      `json:"is_submenu"`
	ParentName string    `json:"parent_name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type MenuMasterCreateResponseRequest struct {
	Id         int       `json:"id"`
	MenuName   string    `validate:"required" json:"menu_name"`
	TitleMenu  string    `validate:"required" json:"title_menu"`
	Path       string    `validate:"required" json:"path"`
	IsSubmenu  bool      `json:"is_submenu"`
	ParentName string    `json:"parent_name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
