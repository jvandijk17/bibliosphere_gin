-- +goose Up
-- Attempt to drop the foreign key. Ignore the error if the foreign key doesn't exist.
ALTER TABLE
    books DROP FOREIGN KEY FK_CBE5A331FE2541D7;

-- Make the library_id column NOT NULL
ALTER TABLE
    books CHANGE library_id library_id INT NOT NULL;

-- Re-add the foreign key constraint
ALTER TABLE
    books
ADD
    CONSTRAINT FK_CBE5A331FE2541D7 FOREIGN KEY (library_id) REFERENCES libraries(id);

-- +goose Down
-- Similar logic for the down migration
ALTER TABLE
    books DROP FOREIGN KEY FK_CBE5A331FE2541D7;

ALTER TABLE
    books CHANGE library_id library_id INT DEFAULT NULL;

ALTER TABLE
    books
ADD
    CONSTRAINT FK_CBE5A331FE2541D7 FOREIGN KEY (library_id) REFERENCES libraries(id);