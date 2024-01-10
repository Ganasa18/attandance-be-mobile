package domain

import "time"

type RoleModel struct {
	Id        int
	Rolename  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RoleAccessModel struct {
	Id        int
	RoleId    int
	Create    bool
	Read      bool
	Update    bool
	Delete    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserMenuAccessModel struct {
	Id           int
	MenuId       int
	UserId       int
	RoleAccessId int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserMenuAccessJoin struct {
	UmaID        int
	MenuID       int
	UserID       int
	RoleAccessID int
	MsRoleID     int
	MsUserID     int
	MsMenuID     int
	RoleID       int
	RoleName     string
	MenuName     string
	TitleMenu    string
	Path         string
	IsSubmenu    bool
	ParentName   string
}
