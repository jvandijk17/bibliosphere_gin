-- +goose Up
-- +goose StatementBegin
ALTER TABLE
    loans
ADD
    estimated_return_date DATE DEFAULT NULL;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
ALTER TABLE
    loans DROP COLUMN estimated_return_date;

-- +goose StatementEnd