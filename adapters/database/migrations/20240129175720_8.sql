-- +goose Up
-- +goose StatementBegin
ALTER TABLE
    loan
ADD
    estimated_return_date DATE DEFAULT NULL;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
ALTER TABLE
    loan DROP COLUMN estimated_return_date;

-- +goose StatementEnd