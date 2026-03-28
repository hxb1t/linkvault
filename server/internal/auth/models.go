package auth

import "time"

type UserEntity struct {
	Id           int       `json:"id"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	RoleId       int       `json:"role_id"`
	PermissionId int       `json:"permission_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username       string `json:"username"`
	HashedPassword string `json:"password"`
}
