package main

import (
	"fmt"
	"net/http"
	"context"
	"github.com/go-chi/chi/v5"
	"strconv"
)

func (apiCfg *apiConfig)handlerGetEntryById(w http.ResponseWriter,r *http.Request){
	idstr := chi.URLParam(r,"id")
	idval,_ := strconv.ParseInt(idstr,10,32)
	id := int32(idval)
	ctx := context.Background()
	user,err := apiCfg.DB.GetEntryFromId(ctx,id)
	if err != nil{
		respondWithError(w,400,fmt.Sprintf("User doesn't exists : %v",err))
		return
	}

	respondWithJSON(w,200,dbuserToUser(user))
}