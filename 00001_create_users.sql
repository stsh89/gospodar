-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE users (email varchar(255), password varchar(255));

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE users;
