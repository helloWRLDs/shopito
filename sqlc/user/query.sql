-- name: GetUserById :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users;

-- name: CreateUser :one
INSERT INTO users(name, email, password)
VALUES($1, $2, $3)
RETURNING id;

-- name: DeleteUser :exec
DELETE FROM users WHERE id=$1;

-- name: UpdateUser :exec
UPDATE users
SET name = $1, email = $2, password = $3, is_admin = $4, is_verified = $5, updated_at = CURRENT_TIMESTAMP
WHERE id = $6;

-- name: IsExistByID :one
SELECT EXISTS(SELECT TRUE from users WHERE id = $1);

-- name: IsExistByEmail :one
SELECT EXISTS(SELECT TRUE FROM users WHERE email = $1);

-- name: IsVerifiedByEmail :one
SELECT is_verified FROM users WHERE email = $1;

-- name: IsVerifiedByID :one
SELECT is_verified FROM users WHERE id = $1;