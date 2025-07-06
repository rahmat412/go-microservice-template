-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users
(
    id         SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name  VARCHAR(255),
    birth_date DATE,
    is_active  BOOLEAN   DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user;
-- +goose StatementEnd`