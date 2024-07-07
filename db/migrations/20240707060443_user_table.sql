-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255),
    surname VARCHAR(255),
    patronymic VARCHAR(255),
    address VARCHAR(255),
    passport VARCHAR(255) NOT NULL,
    passport_series VARCHAR(20) NOT NULL,
    passport_number VARCHAR(20) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (passport)
);

CREATE INDEX idx_passport ON users(passport);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
