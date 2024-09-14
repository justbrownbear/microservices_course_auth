package user_repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// UserRepository defines the methods that any implementation of a user repository
// should provide. This includes creating, retrieving, updating, and deleting users,
// as well as associating the repository with a database transaction.
//
// Methods:
//
//   - CreateUser(ctx context.Context, arg CreateUserParams) (int64, error):
//     Creates a new user with the provided parameters and returns the ID of the created user.
//
//   - GetUser(ctx context.Context, id int64) (GetUserRow, error):
//     Retrieves a user by their ID and returns the user details.
//
//   - UpdateUser(ctx context.Context, arg UpdateUserParams) error:
//     Updates an existing user with the provided parameters.
//
//   - DeleteUser(ctx context.Context, id int64) error:
//     Deletes a user by their ID.
//
//   - WithTx(tx pgx.Tx) *Queries:
//     Associates the repository with a database transaction.
type UserRepository interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (int64, error)
	GetUser(ctx context.Context, id int64) (GetUserRow, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
	DeleteUser(ctx context.Context, id int64) error

	WithTx(tx pgx.Tx) *Queries
}
