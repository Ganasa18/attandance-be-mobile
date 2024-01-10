package domain

import "time"

type MenuModel struct {
	Id         int
	MenuName   string
	TitleMenu  string
	Path       string
	IsSubmenu  bool
	ParentName string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
