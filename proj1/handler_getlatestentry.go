package main

import (
	"fmt"
	"net/http"
	"context"
)

func (apiCfg *apiConfig)handlerGetLatestEntry(w http.ResponseWriter,r *http.Request){
	ctx := context.Background()

	user , err := apiCfg.DB.GetLatestEntry(ctx)
	if err != nil{
		respondWithError(w,400,fmt.Sprintf("Couldnt get user : %v",err))
		return
	}
	respondWithJSON(w,200,dbuserToUser(user))
}