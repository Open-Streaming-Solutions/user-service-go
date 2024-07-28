// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    id, nickname, email
) VALUES (
             $1, $2, $3
         )
RETURNING id, nickname, email
`

type CreateUserParams struct {
	ID       pgtype.UUID
	Nickname string
	Email    string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.ID, arg.Nickname, arg.Email)
	var i User
	err := row.Scan(&i.ID, &i.Nickname, &i.Email)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, nickname, email FROM users
WHERE nickname = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, nickname string) (User, error) {
	row := q.db.QueryRow(ctx, getUser, nickname)
	var i User
	err := row.Scan(&i.ID, &i.Nickname, &i.Email)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, nickname, email FROM users
ORDER BY nickname
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(&i.ID, &i.Nickname, &i.Email); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
