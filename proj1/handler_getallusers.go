package main

import (
	"fmt"
	"net/http"

	"github.com/SandeepSinghSethi/mygoproj/internal/database"
)

func (apicfg *apiConfig) getAllusers(w http.ResponseWriter, r *http.Request, user database.User) {
	// apikey, err := auth.GetAPIKey(r.Header)
	// if err != nil {
	// 	respondWithError(w, 400, fmt.Sprintf("invalid auth request %v", err))
	// 	return
	// }

	// user, err := apicfg.DB.GetUserByAPI(r.Context(), apikey)
	// if err != nil {
	// 	respondWithError(w, 400, fmt.Sprintf("invalid api request %v", err))
	// 	return
	// }
	// above code is handled by middleware

	if user.Username != "sandeep" {
		respondWithError(w, 400, "not enough priviledge ")
		return
	}

	allusers, err := apicfg.DB.GetAllusers(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("unknown sql error %v", err))
		return
	}

	result := ""
	for _, i := range allusers {
		result += i.Username + " " + i.ApiKey + "\n"
	}
	w.Write([]byte(result))
}
