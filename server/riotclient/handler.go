package lol

import (
	"encoding/json"
	"net/http"

	"github.com/jonathanjchung/lol-stats/server/lol"
)

type PlayerInfo struct {
	Account  lol.AccountDto  `json:"account"`
	Summoner lol.SummonerDto `json:"summoner"`
	Matches  []lol.MatchDto  `json:"matches"`
}

func (rc *RiotClient) GetSummonerMatchHistory(w http.ResponseWriter, r *http.Request) {
	summonerName := r.URL.Query().Get("summonerName")
	tagLine := r.URL.Query().Get("tag")

	acc, status, err := lol.GetAccountInfo(summonerName, tagLine, rc.apiKey)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	summoner, status, err := lol.GetSummonerInfo(acc.Puuid, rc.apiKey)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	matches, status, err := lol.GetMatchHistory(acc.Puuid, rc.apiKey)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	var matchHistory []lol.MatchDto
	for i := range matches {
		matchData, status, err := lol.GetMatchDataFromId(matches[i], rc.apiKey)
		if err != nil {
			http.Error(w, err.Error(), status)
			return
		}
		matchHistory = append(matchHistory, matchData)
	}

	result := PlayerInfo{
		Account:  acc, // I don't think you used this yet
		Summoner: summoner,
		Matches:  matchHistory,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
