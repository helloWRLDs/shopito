-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    img_url VARCHAR(255),
    price INT,
    stock INT,
    category_id INT DEFAULT 0 REFERENCES categories(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE products;
-- +goose StatementEnd
