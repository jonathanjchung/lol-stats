package lol

import (
	"fmt"
	"log"
	"net/http"
)

type RiotClient struct {
	apiKey string
}

func NewRiotClient(apiKey string) (*RiotClient, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("API Key is missing")
	}
	return &RiotClient{apiKey: apiKey}, nil
}

func (rc *RiotClient) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	switch r.Method {
	case http.MethodGet:
		rc.HandleGetRequests(w, r)
		return
	default:
		log.Println("default")
		return
	}
}

func (rc *RiotClient) HandleGetRequests(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/summoners":
		rc.GetSummonerMatchHistory(w, r)
		return
	default:
		log.Println("default")
		return
	}
}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*") // this star might be bad but not sure why yet
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
}
