-- +goose Up
-- +goose StatementBegin
CREATE UNIQUE INDEX UNIQ_8D93D649E7927C74 ON users (email);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP INDEX UNIQ_8D93D649E7927C74 ON users;

-- +goose StatementEnd