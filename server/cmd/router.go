package main

import (
	"log"
	"net/http"

	"github.com/jonathanjchung/lol-stats/server/config"
	"github.com/jonathanjchung/lol-stats/server/lol"
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

	mux := http.NewServeMux()
	mux.Handle("/summoners", riotClient)
	http.ListenAndServe(":8080", mux)
}
