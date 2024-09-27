package main

import (
	"fmt"
	"net/http"

	"github.com/SandeepSinghSethi/mygoproj/internal/auth"
	"github.com/SandeepSinghSethi/mygoproj/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apicfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("invalid auth request %v", err))
			return
		}
		user, err := apicfg.DB.GetUserByAPI(r.Context(), apikey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldnt get user %v", err))
			return
		}

		handler(w, r, user)
	}
}
