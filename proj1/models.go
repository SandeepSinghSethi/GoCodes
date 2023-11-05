package main

import (
	"time"
	"github.com/SandeepSinghSethi/mygoproj/internal/database"
)

type User struct {
	ID         int32 `json:"id"`
	CreatedAt  time.Time `json:"created_at`
	ModifiedAt time.Time `json:"modified_at"`
	Name       string `json:"name"`
}

func dbuserToUser(dbuser database.User) User {
	return User{
		ID: dbuser.ID,
		CreatedAt: dbuser.CreatedAt,
		ModifiedAt: dbuser.ModifiedAt,
		Name: dbuser.Name,
	}
}