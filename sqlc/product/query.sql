-- name: ListCategories :many
SELECT * FROM categories;

-- name: CreateCategory :one
INSERT INTO categories(name) VALUES ($1)
RETURNING id;

