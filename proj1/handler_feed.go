package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/SandeepSinghSethi/mygoproj/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type paramaters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	param := paramaters{}
	err := decoder.Decode(&param)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing json %v", err))
		return
	}

	feeds, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:         uuid.New(),
		CreatedAt:  time.Now().UTC(),
		ModifiedAt: time.Now().UTC(),
		Name:       param.Name,
		Url:        param.URL,
		UserID:     user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldnt create feed %v", err))
		return
	}

	respondWithJSON(w, 200, dbfeedToFeed(feeds))
}

func (apicfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apicfg.DB.GetFeed(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feeds : %v", err))
	}

	respondWithJSON(w, 200, dbFeedsslicetoFeeds(feeds))
}
