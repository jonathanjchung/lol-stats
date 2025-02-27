package lol

import (
	"fmt"
	"net/http"

	"github.com/jonathanjchung/lol-stats/server/internal/database"
)

type RiotClient struct {
	apiKey string
	dbConn *database.PostgresInfo
}

func NewRiotClient(apiKey string, pg *database.PostgresInfo) (*RiotClient, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("API Key is missing")
	}

	if pg == nil {
		return nil, fmt.Errorf("database is missing")
	}

	return &RiotClient{apiKey: apiKey, dbConn: pg}, nil
}

func (rc *RiotClient) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/summoners":
		rc.GetSummonerMatchHistory(w, r)
		return
	default:
		return
	}
}
