-- +goose Up
ALTER TABLE
    loans
ADD
    COLUMN book_if_not_returned INT AS (
        CASE
            WHEN return_date IS NULL THEN book_id
            ELSE NULL
        END
    ) VIRTUAL;

ALTER TABLE
    loans
ADD
    UNIQUE INDEX idx_unique_loan_book (book_if_not_returned);

-- +goose Down
ALTER TABLE
    loans DROP INDEX idx_unique_loan_book;

ALTER TABLE
    loans DROP COLUMN book_if_not_returned;