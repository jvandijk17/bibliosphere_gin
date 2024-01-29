-- +goose Up
-- +goose StatementBegin
ALTER TABLE
    user
ADD
    roles LONGTEXT NOT NULL COMMENT '(DC2Type:json)';

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
ALTER TABLE
    user DROP roles;

-- +goose StatementEnd