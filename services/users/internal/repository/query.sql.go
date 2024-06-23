// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package repository

import (
	"context"
	"shopito/services/users/internal/domain"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users(name, email, password)
VALUES($1, $2, $3)
RETURNING id
`

type CreateUserParams struct {
	Name     string
	Email    string
	Password string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Name, arg.Email, arg.Password)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id=$1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, name, email, password, is_admin, is_verified, created_at, updated_at FROM users WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i domain.User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.IsAdmin,
		&i.IsVerified,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, name, email, password, is_admin, is_verified, created_at, updated_at FROM users WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserById(ctx context.Context, id int64) (domain.User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i domain.User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.IsAdmin,
		&i.IsVerified,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const isExistByEmail = `-- name: IsExistByEmail :one
SELECT EXISTS(SELECT TRUE FROM users WHERE email = $1)
`

func (q *Queries) IsExistByEmail(ctx context.Context, email string) (bool, error) {
	row := q.db.QueryRowContext(ctx, isExistByEmail, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const isExistByID = `-- name: IsExistByID :one
SELECT EXISTS(SELECT TRUE from users WHERE id = $1)
`

func (q *Queries) IsExistByID(ctx context.Context, id int64) (bool, error) {
	row := q.db.QueryRowContext(ctx, isExistByID, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const isVerifiedByEmail = `-- name: IsVerifiedByEmail :one
SELECT is_verified FROM users WHERE email = $1
`

func (q *Queries) IsVerifiedByEmail(ctx context.Context, email string) (bool, error) {
	row := q.db.QueryRowContext(ctx, isVerifiedByEmail, email)
	var is_verified bool
	err := row.Scan(&is_verified)
	return is_verified, err
}

const isVerifiedByID = `-- name: IsVerifiedByID :one
SELECT is_verified FROM users WHERE id = $1
`

func (q *Queries) IsVerifiedByID(ctx context.Context, id int64) (bool, error) {
	row := q.db.QueryRowContext(ctx, isVerifiedByID, id)
	var is_verified bool
	err := row.Scan(&is_verified)
	return is_verified, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, name, email, password, is_admin, is_verified, created_at, updated_at FROM users
`

func (q *Queries) ListUsers(ctx context.Context) ([]domain.User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []domain.User
	for rows.Next() {
		var i domain.User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Password,
			&i.IsAdmin,
			&i.IsVerified,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET name = $1, email = $2, password = $3, is_admin = $4, is_verified = $5, updated_at = CURRENT_TIMESTAMP
WHERE id = $6
`

type UpdateUserParams struct {
	Name       string
	Email      string
	Password   string
	IsAdmin    bool
	IsVerified bool
	ID         int64
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.IsAdmin,
		arg.IsVerified,
		arg.ID,
	)
	return err
}
