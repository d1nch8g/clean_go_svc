-- name: InsertUser :one
INSERT INTO users (name, age, description)
VALUES ($1, $2, $3)
RETURNING id;
-- name: SelectUsers :many
SELECT id,
    name,
    age,
    description
FROM users;
-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
-- name: UpdateUser :exec
UPDATE users
SET name = $2,
    age = $3,
    description = $4
WHERE id = $1;