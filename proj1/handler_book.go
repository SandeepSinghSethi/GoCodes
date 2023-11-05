package main

import (
	"net/http"
	"github.com/go-chi/chi/v5"
)

func handlerBook(w http.ResponseWriter,r *http.Request){
	respondWithBook(w,200,chi.URLParam(r,"author"),chi.URLParam(r,"page"),"Book Info")
}