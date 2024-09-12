package user_model

import (
	"database/sql"
	"time"
)


type Role int

const (
    Unknown Role = iota
    User
    Admin
)



type CreateUserRequest struct {
	Name string
	Email string
	Password string
	PasswordConfirm string
	Role Role
}


type GetUserResponse struct {
	ID uint64
	Name string
	Email string
	Role Role
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}


type UpdateUserRequest struct {
	ID uint64
	Name string
	Email string
	Role Role
}
