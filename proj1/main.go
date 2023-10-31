package main

import (
	"os"
	"log"
	"net/http"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	// "github.com/go-chi/jsonp"
	// "github.com/go-chi/render"
	"github.com/go-chi/cors"
)

// type supername struct{
// 	Name string `json:"name"`
// }

func main(){
	godotenv.Load(".env")
	portstring := os.Getenv("PORT")
	if portstring == "" {
		log.Fatal("PORT is not the found in the env file")
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// r.Use(jsonp.Handler)

	// r.Get("/",func(w http.ResponseWriter, r *http.Request){
	// 	// w.Write([]byte("hello world"))
	// 	data := &supername{"waduhekk"}
	// 	render.JSON(w,r,data)
	// })

	// log.Println("Listening on PORT : "+portstring)
	// err := http.ListenAndServe(":"+portstring,r)	
	// if err != nil{
	// 	log.Fatal(err)
	// }

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*","http://*"},
		AllowedMethods: []string{"GET","POST","PUT","DELETE","OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 	300,
	}))

	v1router := chi.NewRouter()
	v1router.Get("/ready",handler)
	v1router.Get("/err",handlerErr)
	v1router.Get("/books/{author}/page/{page}",handlerBook)
	r.Mount("/v1",v1router)


	srv := &http.Server{
		Handler : r,
		Addr : ":"+portstring,
	}
	log.Println("Listening on PORT :"+portstring)
	err := srv.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}
}
