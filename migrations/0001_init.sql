-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(120) NOT NULL CHECK (name <> ''),
    age INTEGER NOT NULL
);
-- +goose Down
DROP TABLE user;