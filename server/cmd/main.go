package main

import (
	"log"
	"net/http"

	"github.com/jonathanjchung/lol-stats/server/config"
	"github.com/jonathanjchung/lol-stats/server/internal/database"
	"github.com/jonathanjchung/lol-stats/server/internal/middleware"
	lol "github.com/jonathanjchung/lol-stats/server/internal/riotclient"
	_ "github.com/lib/pq"
)

func main() {
	a := config.GetEnvVariable("riot_api_key")
	u := config.GetEnvVariable("user")
	p := config.GetEnvVariable("password")
	d := config.GetEnvVariable("database")
	h := config.GetEnvVariable("host")
	if a == "" || u == "" || p == "" || d == "" || h == "" {
		log.Fatal(".env configs not set")
	}

	// user, password, database, host
	db, err := database.NewPostgresConnection(u, p, d, h)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	riotClient, err := lol.NewRiotClient(a, db)
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
