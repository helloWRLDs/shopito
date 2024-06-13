-- name: ListCategories :many
SELECT * FROM categories;

-- name: CreateCategory :one
INSERT INTO categories(name) VALUES ($1)
RETURNING id;

-- name: IsCategoryExistByID :one
SELECT EXISTS(SELECT TRUE FROM categories WHERE id = $1);

-- name: IsCategoryExistByName :one
SELECT EXISTS(SELECT TRUE FROM categories WHERE name = $1);

-- name: IsProductExistByID :one
SELECT EXISTS(SELECT TRUE FROM products WHERE id = $1);

-- name: IsProductExistByName :one
SELECT EXISTS(SELECT TRUE FROM products WHERE name = $1);

-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = $1;

-- name: UpdateCategory :exec
UPDATE categories SET name = $1 WHERE id = $2;

-- name: GetCategoryById :one
SELECT * FROM categories WHERE id = $1;

-- name: GetCategoryByName :one
SELECT * FROM categories WHERE name = $1;

-- name: GetProductById :one
SELECT * FROM products WHERE id = $1;

-- name: CreateProduct :one
INSERT INTO products(name, price, stock, category_id, img_url)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;

-- name: UpdateProduct :exec
UPDATE products
SET name = $1, img_url = $2, price = $3, stock = $4, category_id = $5
WHERE id = $6;

-- name: DeleteProduct :exec
DELETE FROM products WHERE id = $1;

-- name: ListProducts :many
SELECT * FROM products
WHERE name ILIKE $1
ORDER BY $2::text ASC
LIMIT $3 OFFSET $4;

-- name: ListProductsByCategoryName :many
SELECT * FROM products as p 
JOIN categories as c 
ON p.category_id = c.id 
WHERE c.name ILIKE $1
LIMIT $2 OFFSET $3;