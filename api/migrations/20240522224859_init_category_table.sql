-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS categories(
    id SERIAL PRIMARY KEY,
    name VARCHAR(60) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE categories;
-- +goose StatementEnd
