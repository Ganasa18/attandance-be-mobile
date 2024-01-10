package web

import "time"

type RoleMasterResponseRequest struct {
	Id        int       `json:"id"`
	Rolename  string    `json:"role_name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type RoleMasterCreateResponseRequest struct {
	Id        int       `json:"id"`
	Rolename  string    `validate:"required,min=4" json:"role_name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type RoleAccessMasterResponseRequest struct {
	Id        int       `json:"id"`
	RoleId    int       `json:"role_id"`
	Create    bool      `json:"create"`
	Read      bool      `json:"read"`
	Update    bool      `json:"update"`
	Delete    bool      `json:"delete"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type RoleAccessMasterCreateResponseRequest struct {
	Id        int       `json:"id"`
	RoleId    int       `validate:"required" json:"role_id"`
	Create    bool      `json:"create"`
	Read      bool      `json:"read"`
	Update    bool      `json:"update"`
	Delete    bool      `json:"delete"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
