package web

import "time"

type UserResponseRequest struct {
	Id           int    `json:"id"`
	UserUniqueId string `json:"user_unique_id"`
	Username     string `json:"name"`
	Email        string `json:"email"`
	// Password     *string    `json:"password"`
	Token    *string `json:"token"`
	IsActive bool    `json:"is_active"`
	// UserRole  *int       `json:"user_role"`
	// RoleId    *int       `json:"role_id"`
	RoleName  *string    `json:"role_name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type UserCreateRequest struct {
	Id           int        `json:"id"`
	UserUniqueId string     `json:"user_unique_id"`
	Username     string     `validate:"required,max=200,min=1" json:"username"`
	Email        string     `validate:"required,email" json:"email"`
	Password     string     `validate:"required,max=200,min=1" json:"password"`
	IsActive     bool       `json:"is_active"`
	UserRole     *int       `json:"user_role"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}
