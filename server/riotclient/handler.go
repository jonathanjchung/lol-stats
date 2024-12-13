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

	acc := lol.GetAccountInfo(summonerName, tagLine, rc.apiKey)
	summoner := lol.GetSummonerInfo(acc.Puuid, rc.apiKey)

	matches := lol.GetMatchHistory(acc.Puuid, rc.apiKey)

	var matchHistory []lol.MatchDto
	for i := range matches {
		matchData := lol.GetMatchDataFromId(matches[i], rc.apiKey)
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
