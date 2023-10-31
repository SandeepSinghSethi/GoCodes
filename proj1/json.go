package main

import (
	"net/http"
	"encoding/json"
	"log"
)

func respondWithBook(w http.ResponseWriter,code int,author , page , msg string){
	if code != 200{
		log.Println("Something went wrong")
	}

	type bookResp struct{
		Message string `json:"message"`
		Author string `json:"auth"`
		Page string `json:"pg"`
	}

	respondWithJSON(w,code,bookResp{
		Message : msg,
		Author : author,
		Page : page,
	})
}

func respondWithError(w http.ResponseWriter,code int,msg string){
	if code > 499 {
		log.Println("Server Responding with 5xx Error Codes")
	}
	type errResp struct{
		Error string `json:"error"`
	}

	respondWithJSON(w,code,errResp{
		Error : msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int , payload interface{}){
	data,err := json.Marshal(payload)
	if err != nil{
		log.Println("Failed to marshal response %v",payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(code)
	w.Write(data)
}