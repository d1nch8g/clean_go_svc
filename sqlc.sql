-- name: InsertUser :one
INSERT INTO users (name, age)
VALUES ($1, $2)
RETURNING id;
-- name: SelectUser :one
SELECT name,
    age
FROM users
WHERE id = $1;
-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
-- name: UpdateUser :exec
UPDATE users
SET name = $2,
    age = $3
WHERE id = $1;