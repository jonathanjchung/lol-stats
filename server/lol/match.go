package lol

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jonathanjchung/lol-stats/server/config"
)

func GetMatchHistory(puuid string, apiKey string) ([]string, int, error) {
	url := fmt.Sprintf(config.MatchHistoryEndpoint, puuid, apiKey)

	if puuid == "" {
		return nil, http.StatusBadRequest, fmt.Errorf("missing puuid")
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("failed to make HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, resp.StatusCode, fmt.Errorf("received non-OK status code: %d", resp.StatusCode)
	}

	var matchList []string
	if err := json.NewDecoder(resp.Body).Decode(&matchList); err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("failed to decode response: %v", err)
	}

	return matchList, http.StatusOK, nil
}

func GetMatchDataFromId(matchId string, apiKey string) (MatchDto, int, error) {
	url := fmt.Sprintf(config.MatchDataEndpoint, matchId, apiKey)

	if matchId == "" {
		return MatchDto{}, http.StatusBadRequest, fmt.Errorf("missing matchId")
	}

	resp, err := http.Get(url)
	if err != nil {
		return MatchDto{}, http.StatusInternalServerError, fmt.Errorf("failed to make HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return MatchDto{}, resp.StatusCode, fmt.Errorf("received non-OK status code: %d", resp.StatusCode)
	}

	var matchData MatchDto
	if err := json.NewDecoder(resp.Body).Decode(&matchData); err != nil {
		return MatchDto{}, http.StatusInternalServerError, fmt.Errorf("failed to decode response: %v", err)
	}

	return matchData, http.StatusOK, nil
}
