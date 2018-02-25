-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE users;
