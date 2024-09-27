-- +goose Up

CREATE TABLE users (
	id UUID PRIMARY KEY,
	created_at TIMESTAMP NOT NULL,
	modified_at TIMESTAMP NOT NULL,
	username VARCHAR(64) UNIQUE NOT NULL
);

-- +goose Down

DROP TABLE users;