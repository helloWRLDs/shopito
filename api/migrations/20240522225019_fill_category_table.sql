-- +goose Up
-- +goose StatementBegin
INSERT INTO categories(name) VALUES('T-shirt');
INSERT INTO categories(name) VALUES('Pants');
INSERT INTO categories(name) VALUES('Accessories');
INSERT INTO categories(name) VALUES('Bags');
INSERT INTO categories(name) VALUES('Shoes');
INSERT INTO categories(name) VALUES('Outerwear');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM categories;
-- +goose StatementEnd
