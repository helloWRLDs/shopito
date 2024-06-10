-- name: GetCategory :one
SELECT * FROM categories WHERE id = $1;

-- name: ListCategories :many
SELECT * FROM categories;

-- name: CreateCategory :one
INSERT INTO categories(name) VALUES($1)
RETURNING id;

-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = $1;

-- name: UpdateCategory :exec
UPDATE categories SET name = $1;