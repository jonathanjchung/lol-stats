package lol

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jonathanjchung/lol-stats/server/config"
)

func GetMatchHistory(puuid string, apiKey string) []string {
	url := fmt.Sprintf(config.MatchHistoryEndpoint, puuid, apiKey)

	if puuid == "" {
		log.Println("No puuid")
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	var matchList []string
	if err := json.NewDecoder(resp.Body).Decode(&matchList); err != nil {
		log.Println(err)
		return nil
	}

	return matchList
}

func GetMatchDataFromId(matchId string, apiKey string) MatchDto {
	url := fmt.Sprintf(config.MatchDataEndpoint, matchId, apiKey)

	if matchId == "" {
		log.Println("No matchId")
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	var matchData MatchDto
	if err := json.NewDecoder(resp.Body).Decode(&matchData); err != nil {
		log.Println(err)
		return MatchDto{}
	}

	return matchData
}
