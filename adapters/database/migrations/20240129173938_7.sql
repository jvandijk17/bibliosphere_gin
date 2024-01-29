-- +goose Up
ALTER TABLE
    loan
ADD
    COLUMN book_if_not_returned INT AS (
        CASE
            WHEN return_date IS NULL THEN book_id
            ELSE NULL
        END
    ) VIRTUAL;

ALTER TABLE
    loan
ADD
    UNIQUE INDEX idx_unique_loan_book (book_if_not_returned);

-- +goose Down
ALTER TABLE
    loan DROP INDEX idx_unique_loan_book;

ALTER TABLE
    loan DROP COLUMN book_if_not_returned;