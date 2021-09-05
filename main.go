package main

import (
	"log"
	"net/http"
	"os"

	"alexis-api.com/main/handlers"
	"alexis-api.com/main/server"
)

var (
	apiPort = os.Getenv("API_PORT")
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/anime-quotes", handlers.AnimeQuotes)

	srv := server.New(mux, apiPort)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("server failed to start: %v", err)
	}

}
