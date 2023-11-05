package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"time"
	"context"
	"github.com/SandeepSinghSethi/mygoproj/internal/database"
)

func (apiCfg *apiConfig)handlerCreateUser(w http.ResponseWriter,r *http.Request){
	type paramaters struct{
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	param := paramaters{}
	err := decoder.Decode(&param)

	if err != nil{
		respondWithError(w,400,fmt.Sprintf("Error parsing json %v",err))
		return
	}

	ctx := context.Background()

	user,err := apiCfg.DB.CreateUser(ctx,database.CreateUserParams{
		CreatedAt: time.Now().UTC(),
		ModifiedAt: time.Now().UTC(),
		Name: param.Name,
	})
	if err != nil{
		respondWithError(w,400,fmt.Sprintf("Couldnt create user %v",err))
		return
	}

	respondWithJSON(w,200,user)
}