package web

import "time"

type AuthLoginResponseRequest struct {
	Id           int       `json:"id"`
	UserUniqueId string    `json:"user_unique_id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Token        *string   `json:"token"`
	IsActive     bool      `json:"isActive"`
	UserRole     *int      `json:"user_role"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type AuthLoginRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}
