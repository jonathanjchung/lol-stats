package lol

import (
	"fmt"
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
	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/summoners":
		rc.GetSummonerMatchHistory(w, r)
		return
	default:
		return
	}
}
