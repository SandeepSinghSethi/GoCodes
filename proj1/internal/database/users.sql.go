// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: users.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const createUser = `-- name: CreateUser :execresult
INSERT INTO users (created_at, modified_at , name)
VALUES (?,?,?)
`

type CreateUserParams struct {
	CreatedAt  time.Time
	ModifiedAt time.Time
	Name       string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser, arg.CreatedAt, arg.ModifiedAt, arg.Name)
}

const getEntryFromId = `-- name: GetEntryFromId :one
SELECT id, created_at, modified_at, name FROM users WHERE id = ?
`

func (q *Queries) GetEntryFromId(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getEntryFromId, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.ModifiedAt,
		&i.Name,
	)
	return i, err
}

const getLatestEntry = `-- name: GetLatestEntry :one
SELECT id, created_at, modified_at, name FROM users WHERE id = LAST_INSERT_ID()
`

func (q *Queries) GetLatestEntry(ctx context.Context) (User, error) {
	row := q.db.QueryRowContext(ctx, getLatestEntry)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.ModifiedAt,
		&i.Name,
	)
	return i, err
}
