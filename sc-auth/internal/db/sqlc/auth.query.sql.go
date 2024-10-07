// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: auth.query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO auth (username, email, password)
VALUES ($1, $2, $3)
    RETURNING id, username, email, password
`

type CreateUserParams struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRow struct {
	ID       pgtype.UUID `json:"id"`
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Password string      `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Username, arg.Email, arg.Password)
	var i CreateUserRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM auth
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, username, email, password
FROM auth
WHERE email = $1
    LIMIT 1
`

type GetUserByEmailRow struct {
	ID       pgtype.UUID `json:"id"`
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Password string      `json:"password"`
}

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (GetUserByEmailRow, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i GetUserByEmailRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, username, email, password, last_login
FROM auth
WHERE id = $1
    LIMIT 1
`

func (q *Queries) GetUserById(ctx context.Context, id pgtype.UUID) (Auth, error) {
	row := q.db.QueryRow(ctx, getUserById, id)
	var i Auth
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.LastLogin,
	)
	return i, err
}

const updateLastLogin = `-- name: UpdateLastLogin :exec
UPDATE auth
SET last_login = NOW()
WHERE id = $1
`

func (q *Queries) UpdateLastLogin(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, updateLastLogin, id)
	return err
}

const updateUserPassword = `-- name: UpdateUserPassword :exec
UPDATE auth
SET password = $1, updated_at = NOW()
WHERE id = $2
`

type UpdateUserPasswordParams struct {
	NewPassword string      `json:"new_password"`
	ID          pgtype.UUID `json:"id"`
}

func (q *Queries) UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error {
	_, err := q.db.Exec(ctx, updateUserPassword, arg.NewPassword, arg.ID)
	return err
}
