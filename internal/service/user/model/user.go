package user_model

import (
	"database/sql"
	"time"
)

// Role represents the role of a user in the system.
// It is defined as an integer type where each value corresponds to a specific role.
type Role int

const (
	// Unknown represents an undefined or unrecognized role.
	Unknown Role = iota
	// User represents an user role.
	User
	// Admin represents an admin role.
	Admin
)

// CreateUserRequest represents the data required to create a new user.
// It includes the user's name, email, password, password confirmation, and role.
type CreateUserRequest struct {
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	Role            Role
}

// GetUserResponse represents the response structure for retrieving user information.
// It includes the user's ID, name, email, role, and timestamps for creation and updates.
type GetUserResponse struct {
	ID        uint64       `redis:"id"`
	Name      string       `redis:"name"`
	Email     string       `redis:"email"`
	Role      Role         `redis:"role"`
	CreatedAt time.Time    `redis:"created_at"`
	UpdatedAt sql.NullTime `redis:"updated_at"`
}

// GetUserResponseForRedis represents the structure of a user response
// that is stored in Redis. It includes the user's ID, name, email,
// role, and timestamps for when the user was created and last updated.
type GetUserResponseForRedis struct {
	ID        uint64 `redis:"id"`
	Name      string `redis:"name"`
	Email     string `redis:"email"`
	Role      Role   `redis:"role"`
	CreatedAt int64  `redis:"created_at"`
	UpdatedAt int64  `redis:"updated_at"`
}

// UpdateUserRequest represents the data required to update a user's information.
// It includes the user's ID, name, email, and role.
type UpdateUserRequest struct {
	ID    uint64
	Name  string
	Email string
	Role  Role
}
