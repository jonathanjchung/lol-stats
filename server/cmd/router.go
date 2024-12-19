package main

import (
	"log"
	"net/http"

	"github.com/jonathanjchung/lol-stats/server/config"
	"github.com/jonathanjchung/lol-stats/server/middleware"
	lol "github.com/jonathanjchung/lol-stats/server/riotclient"
)

func main() {
	dotenv := config.GetEnvVariable("riot_api_key")
	if dotenv == "" {
		log.Fatal("Riot_API_Key environment variable not set")
	}

	riotClient, err := lol.NewRiotClient(dotenv)
	if err != nil {
		log.Fatalf("Error creating RiotClient: %v", err)
	}

	commonMiddleware := []middleware.Middleware{
		middleware.LogMiddleware,
		middleware.CORSMiddleware,
	}

	mux := http.NewServeMux()
	mux.Handle("/summoners", middleware.Chain(
		func(w http.ResponseWriter, r *http.Request) {
			riotClient.ServeHTTP(w, r)
		},
		commonMiddleware...,
	))
	http.ListenAndServe(":8080", mux)
}
