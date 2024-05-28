package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/LKarataev/Film-Library-API/internal/service"
	_ "github.com/lib/pq"
)

var (
	DSN = "postgresql://postgres:my_secret_password@FilmLibraryDB:5432?sslmode=disable"
)

type Response struct {
	Body  interface{} `json:"response,omitempty"`
	Error string      `json:"error,omitempty"`
}

func main() {
	log.SetOutput(os.Stdout)
	db, err := sql.Open("postgres", DSN)
	if err != nil {
		log.Println("SQL connection error: ", err)
		return
	}

	err = db.Ping()
	if err != nil {
		log.Println("Ping to DB error: ", err)
		return
	}

	api, err := service.NewFilmLibraryApi(db)
	if err != nil {
		log.Println("NewFilmLibraryApi error: ", err)
		return
	}

	muxHandler := api.ConfigureRouter()

	log.Println("starting server at :8080")
	http.ListenAndServe(":8080", muxHandler)
}
