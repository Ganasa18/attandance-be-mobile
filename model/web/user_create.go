package web

import "time"

type UserCreateRequest struct {
	Username  string     `validate:"required,max=200,min=1" json:"username"`
	Email     string     `validate:"required,max=200,min=1" json:"email"`
	Password  string     `validate:"required,max=200,min=1" json:"password"`
	IsActive  bool       `validate:"required,max=200,min=1" json:"isActive"`
	CreatedAt time.Time  `validate:"required,max=200,min=1" json:"createdAt"`
	UpdatedAt time.Time  `validate:"required,max=200,min=1" json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
