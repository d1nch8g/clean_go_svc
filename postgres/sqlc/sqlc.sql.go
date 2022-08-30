// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: sqlc.sql

package sqlc

import (
	"context"
)

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const insertUser = `-- name: InsertUser :one
INSERT INTO users (name, age, description)
VALUES ($1, $2, $3)
RETURNING id
`

type InsertUserParams struct {
	Name        string
	Age         int32
	Description string
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (int32, error) {
	row := q.db.QueryRow(ctx, insertUser, arg.Name, arg.Age, arg.Description)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const selectUsers = `-- name: SelectUsers :many
SELECT id,
    name,
    age,
    description
FROM users
`

type SelectUsersRow struct {
	ID          int32
	Name        string
	Age         int32
	Description string
}

func (q *Queries) SelectUsers(ctx context.Context) ([]SelectUsersRow, error) {
	rows, err := q.db.Query(ctx, selectUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SelectUsersRow
	for rows.Next() {
		var i SelectUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Age,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET name = $2,
    age = $3,
    description = $4
WHERE id = $1
`

type UpdateUserParams struct {
	ID          int32
	Name        string
	Age         int32
	Description string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser,
		arg.ID,
		arg.Name,
		arg.Age,
		arg.Description,
	)
	return err
}