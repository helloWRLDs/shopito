-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS verification (
    id SERIAL PRIMARY KEY,
    code VARCHAR(6) NOT NULL,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE verification;
-- +goose StatementEnd
