-- +goose Up
-- +goose StatementBegin
INSERT INTO users(id, name, email, password, is_admin, is_verified) 
VALUES(0, 'admin', 'admin@gmail.com', '$2a$12$I66wAiXk93UMv3zUL60YS.QZHSy.RGS8ZDlQQJN9W4rCqWqkZ2Uru', true, true);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users WHERE id=0;
-- +goose StatementEnd
