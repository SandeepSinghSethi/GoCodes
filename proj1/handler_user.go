package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/SandeepSinghSethi/mygoproj/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type paramaters struct {
		Name string `json:"username"`
	}
	decoder := json.NewDecoder(r.Body)
	param := paramaters{}
	err := decoder.Decode(&param)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing json %v", err))
		return
	}

	ctx := context.Background()

	user, err := apiCfg.DB.CreateUser(ctx, database.CreateUserParams{
		ID:         uuid.New(),
		CreatedAt:  time.Now().UTC(),
		ModifiedAt: time.Now().UTC(),
		Username:   param.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldnt create user %v", err))
		return
	}

	respondWithJSON(w, 200, dbuserToUser(user))
}

// handler for api key auth
func (apicfg *apiConfig) handlerGetUserByAPI(w http.ResponseWriter, r *http.Request, user database.User) {

	respondWithJSON(w, 200, dbuserToUser(user))
}

func (apicfg *apiConfig) handlerGetPostByUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apicfg.DB.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't Fetch Posts by current user . %v", err))
	}

	respondWithJSON(w, 200, databasePostsToPosts(posts))
}
