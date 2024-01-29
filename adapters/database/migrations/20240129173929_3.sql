-- +goose Up
-- +goose StatementBegin
ALTER TABLE
    user CHANGE password_hash password VARCHAR(255) NOT NULL;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
ALTER TABLE
    user CHANGE password password_hash VARCHAR(255) NOT NULL;

-- +goose StatementEnd