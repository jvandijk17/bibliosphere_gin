-- +goose Up
-- Drop foreign key constraints
ALTER TABLE
    loan DROP FOREIGN KEY FK_C5D30D03A76ED395;

ALTER TABLE
    loan DROP FOREIGN KEY FK_C5D30D0316A2B381;

-- Alter the columns
ALTER TABLE
    loan CHANGE user_id user_id INT NOT NULL,
    CHANGE book_id book_id INT NOT NULL;

-- Recreate the foreign key constraints
ALTER TABLE
    loan
ADD
    CONSTRAINT FK_C5D30D03A76ED395 FOREIGN KEY (user_id) REFERENCES user (id);

ALTER TABLE
    loan
ADD
    CONSTRAINT FK_C5D30D0316A2B381 FOREIGN KEY (book_id) REFERENCES book (id);

-- +goose Down
-- Drop foreign key constraints
ALTER TABLE
    loan DROP FOREIGN KEY FK_C5D30D03A76ED395;

ALTER TABLE
    loan DROP FOREIGN KEY FK_C5D30D0316A2B381;

-- Revert the columns
ALTER TABLE
    loan CHANGE user_id user_id INT DEFAULT NULL,
    CHANGE book_id book_id INT DEFAULT NULL;

-- Recreate the foreign key constraints
ALTER TABLE
    loan
ADD
    CONSTRAINT FK_C5D30D03A76ED395 FOREIGN KEY (user_id) REFERENCES user (id);

ALTER TABLE
    loan
ADD
    CONSTRAINT FK_C5D30D0316A2B381 FOREIGN KEY (book_id) REFERENCES book (id);