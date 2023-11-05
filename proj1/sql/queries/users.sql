-- name: CreateUser :execresult
INSERT INTO users (created_at, modified_at , name)
VALUES (?,?,?);

-- name: GetLatestEntry :one
SELECT * FROM users WHERE id = LAST_INSERT_ID();

-- name: GetEntryFromId :one
SELECT * FROM users WHERE id = ?;