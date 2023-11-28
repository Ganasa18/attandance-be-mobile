package domain

import "time"

type UserModel struct {
	Id        int
	Username  string
	Email     string
	Password  string
	Token     *string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
