package domain

import "time"

type UserModel struct {
	Id           int
	UserUniqueId string
	Username     string
	Email        string
	Password     string
	Token        *string
	IsActive     bool
	UserRole     *int
	RoleId       *int
	RoleName     *string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

type UserRegisterModel struct {
	Id           int
	UserUniqueId string
	Username     string
	Email        string
	Password     string
	IsActive     bool
	UserRole     *int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserInfoModel struct {
	Id         int
	UserId     int
	FirstName  string
	LastName   string
	Address    string
	ProfileImg string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type SelfServiceUserModel struct {
	Id           int
	UserId       int
	RequestType  string
	Approval     string
	DateApproval *time.Time
	CreatedAt    time.Time
}
