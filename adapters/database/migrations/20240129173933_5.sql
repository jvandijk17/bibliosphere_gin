-- +goose Up
-- +goose StatementBegin
ALTER TABLE
    users
ADD
    roles LONGTEXT NOT NULL COMMENT '(DC2Type:json)';

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
ALTER TABLE
    users DROP roles;

-- +goose StatementEnd