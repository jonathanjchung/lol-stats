package lol

import (
	"encoding/json"
	"log"
	"net/http"
)

type PlayerInfo struct {
	Summoner SummonerDto `json:"summoner"`
	Matches  []MatchDto  `json:"matches"`
}

func (rc *RiotClient) GetSummonerMatchHistory(w http.ResponseWriter, r *http.Request) {
	summonerName := r.URL.Query().Get("summonerName")
	tagLine := r.URL.Query().Get("tag")

	acc := GetAccountInfo(summonerName, tagLine, rc.apiKey)
	summoner := GetSummonerInfo(acc.Puuid, rc.apiKey)

	matches := GetMatchHistory(acc.Puuid, rc.apiKey)

	var matchHistory []MatchDto
	for i, _ := range matches {
		matchData := GetMatchDataFromId(matches[i], rc.apiKey)
		matchHistory = append(matchHistory, matchData)
	}

	result := PlayerInfo{
		Summoner: summoner,
		Matches:  matchHistory,
	}

	log.Println(result)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
