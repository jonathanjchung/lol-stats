package lol

import (
	"encoding/json"
	"net/http"

	"github.com/jonathanjchung/lol-stats/server/internal/database"
	"github.com/jonathanjchung/lol-stats/server/internal/lol"
)

type PlayerInfo struct {
	Account  database.AccountDto  `json:"account"`
	Summoner database.SummonerDto `json:"summoner"`
	Matches  []database.MatchDto  `json:"matches"`
}

func (rc *RiotClient) GetSummonerMatchHistory(w http.ResponseWriter, r *http.Request) {
	summonerName := r.URL.Query().Get("summonerName")
	tagLine := r.URL.Query().Get("tag")
	var player *database.Player

	if rc.dbConn.SummonerExists(summonerName, tagLine) {
		player = rc.dbConn.GetSummonerInfoFromDB(summonerName, tagLine)
	} else {
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

		var matchHistory []database.MatchDto
		for i := range matches {
			matchData, status, err := lol.GetMatchDataFromId(matches[i], rc.apiKey)
			if err != nil {
				http.Error(w, err.Error(), status)
				return
			}
			matchHistory = append(matchHistory, matchData)
		}

		player = &database.Player{
			Puuid:        acc.Puuid,
			SummonerName: acc.GameName,
			TagLine:      acc.TagLine,
			Account:      acc,
			Summoner:     summoner,
			Matches:      matchHistory,
		}

		rc.dbConn.InsertSummonerInfoIntoDB(player)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(player); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
