package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/SandeepSinghSethi/mygoproj/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type paramaters struct {
		FeedID uuid.UUID `json:"feedid"`
	}
	decoder := json.NewDecoder(r.Body)
	param := paramaters{}
	err := decoder.Decode(&param)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing json %v", err))
		return
	}

	feeds, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:         uuid.New(),
		CreatedAt:  time.Now().UTC(),
		ModifiedAt: time.Now().UTC(),
		UserID:     user.ID,
		FeedID:     param.FeedID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldnt create feed follows %v", err))
		return
	}

	respondWithJSON(w, 200, dbfeedfollowToFeedFollow(feeds))
}

func (apicfg *apiConfig) handlerGetFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	feed_flw, err := apicfg.DB.GetFeedFollow(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feeds for the current user : %v", err))
	}

	respondWithJSON(w, 200, dbfeedfollowSliceToFeedFollow(feed_flw))
}

func (apicfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	feedfollowidparam := chi.URLParam(r, "feedfollowID")
	feedfollowid, err := uuid.Parse(feedfollowidparam)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feed id in the http url  : %v", err))
	}

	err = apicfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedfollowid,
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(w, 404, fmt.Sprintf("Couldn't delete feeds for the user : %v", err))
	}
	respondWithJSON(w, 200, struct{}{})
}
