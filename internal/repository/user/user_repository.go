package user_repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)


type UserRepository interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (int64, error)
	GetUser(ctx context.Context, id int64) (GetUserRow, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
	DeleteUser(ctx context.Context, id int64) error

	WithTx(tx pgx.Tx) *Queries
}
