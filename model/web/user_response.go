package web

import "time"

type UserResponseRequest struct {
	Id        int        `json:"id"`
	Username  string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Token     *string    `json:"token"`
	IsActive  bool       `json:"isActive"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
