package main

import (
	"log"
	"os"
	"time"

	// "fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	// "github.com/go-chi/jsonp"
	// "github.com/go-chi/render"
	"database/sql"

	"github.com/SandeepSinghSethi/mygoproj/internal/database"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
)

// type supername struct{
// 	Name string `json:"name"`
// }

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")
	portstring := os.Getenv("PORT")
	if portstring == "" {
		log.Fatal("PORT is not the found in the env file")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("dbstring is not the found in the env file")
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
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1router := chi.NewRouter()
	v1router.Get("/ready", handler)
	v1router.Get("/err", handlerErr)
	v1router.Get("/books/{author}/page/{page}", handlerBook)

	srv := &http.Server{
		Handler: r,
		Addr:    ":" + portstring,
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err.Error())
	}

	dbs := database.New(db)
	apiCfg := apiConfig{DB: dbs}

	v1router.Post("/users", apiCfg.handlerCreateUser)
	v1router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUserByAPI))
	v1router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1router.Get("/feeds", apiCfg.handlerGetFeeds)
	v1router.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
	v1router.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollow))
	v1router.Delete("/feed_follows/{feedfollowID}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollow))
	v1router.Get("/posts", apiCfg.middlewareAuth(apiCfg.handlerGetPostByUser))

	v1router.Get("/allusers", apiCfg.middlewareAuth(apiCfg.getAllusers))

	r.Mount("/v1", v1router)
	r.Get("/", func(w http.ResponseWriter, resp *http.Request) {
		apiCfg.getRoutes(r, w, resp)
	})
	// func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("helloworld"))
	// })

	log.Println("Listening on PORT :" + portstring)
	go scrapeRss(dbs, 10, time.Minute)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
