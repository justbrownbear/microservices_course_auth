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
	ID			uint64			`redis:"id"`
	Name		string			`redis:"name"`
	Email		string			`redis:"email"`
	Role		Role			`redis:"role"`
	CreatedAt	time.Time		`redis:"created_at"`
	UpdatedAt	sql.NullTime	`redis:"updated_at"`
}


type GetUserResponseForRedis struct {
	ID			uint64		`redis:"id"`
	Name		string		`redis:"name"`
	Email		string		`redis:"email"`
	Role		Role		`redis:"role"`
	CreatedAt	int64		`redis:"created_at"`
	UpdatedAt	int64		`redis:"updated_at"`
}




type UpdateUserRequest struct {
	ID uint64
	Name string
	Email string
	Role Role
}
