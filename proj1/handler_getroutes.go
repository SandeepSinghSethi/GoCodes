package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (apiCfg *apiConfig) getRoutes(router *chi.Mux, w http.ResponseWriter, r *http.Request) {
	var allroutes []string

	chi.Walk(router, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		// allroutes = append(allroutes, fmt.Sprintf("[%s] : %s has %d middlewares \n", method, route, len(middlewares)))
		allroutes = append(allroutes, fmt.Sprintf("[%s] : %s \n", method, route))
		return nil
	})
	// fmt.Println(allroutes)
	result := ""
	for _, str := range allroutes {
		result += str
	}
	// fmt.Println(result)

	w.Write([]byte(result))
	w.Write([]byte("\n\nALL ROUTES FOR THIS WEBSITE ARE ABOVE"))
}
