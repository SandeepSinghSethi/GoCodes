-- +goose Up

CREATE TABLE users (
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	created_at TIMESTAMP NOT NULL,
	modified_at TIMESTAMP NOT NULL,
	name TEXT NOT NULL
);

-- +goose Down

DROP TABLE users;