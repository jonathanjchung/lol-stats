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
