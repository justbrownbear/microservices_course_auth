// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: queries.sql

package user_repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO public.users (name, email, role, password_hash)
	VALUES ($1::text, $2::text, $3::smallint, $4::text)
	RETURNING id
`

type CreateUserParams struct {
	Name         string
	Email        string
	Role         int16
	PasswordHash string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (int64, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Name,
		arg.Email,
		arg.Role,
		arg.PasswordHash,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const deleteUser = `-- name: DeleteUser :exec
UPDATE public.users
	SET
		is_deleted = TRUE,
		delete_timestamp = NOW()
	WHERE
		id = $1
		AND NOT is_deleted
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, name, email, role, create_timestamp, update_timestamp
	FROM public.users
	WHERE
		id = $1
		AND NOT is_deleted
`

type GetUserRow struct {
	ID              int64
	Name            string
	Email           string
	Role            int16
	CreateTimestamp pgtype.Timestamp
	UpdateTimestamp pgtype.Timestamp
}

func (q *Queries) GetUser(ctx context.Context, id int64) (GetUserRow, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i GetUserRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Role,
		&i.CreateTimestamp,
		&i.UpdateTimestamp,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE public.users
	SET
		name = COALESCE( NULLIF( $1::text, '' ), name ),
		email = COALESCE( NULLIF( $2::text, '' ), email ),
		role = COALESCE( NULLIF( $3::smallint, 0 ), role ),
		update_timestamp = NOW()
	WHERE
		id = $4
		AND NOT is_deleted
`

type UpdateUserParams struct {
	Name  string
	Email string
	Role  int16
	ID    int64
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser,
		arg.Name,
		arg.Email,
		arg.Role,
		arg.ID,
	)
	return err
}
