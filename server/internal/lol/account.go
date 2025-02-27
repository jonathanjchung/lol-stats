package lol

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jonathanjchung/lol-stats/server/config"
	"github.com/jonathanjchung/lol-stats/server/internal/database"
)

func GetAccountInfo(summonerName string, tagLine string, apiKey string) (database.AccountDto, int, error) {
	url := fmt.Sprintf(config.AccountInfoEndpoint, summonerName, tagLine, apiKey)

	if summonerName == "" || tagLine == "" {
		return database.AccountDto{}, http.StatusBadRequest, fmt.Errorf("missing summonerName or tagLine")
	}

	resp, err := http.Get(url)
	if err != nil {
		return database.AccountDto{}, http.StatusInternalServerError, fmt.Errorf("failed to make HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return database.AccountDto{}, resp.StatusCode, fmt.Errorf("received non-OK status code: %d", resp.StatusCode)
	}

	var acc database.AccountDto
	if err := json.NewDecoder(resp.Body).Decode(&acc); err != nil {
		return database.AccountDto{}, http.StatusInternalServerError, fmt.Errorf("failed to decode response: %v", err)
	}

	return acc, http.StatusOK, nil
}

func GetSummonerInfo(encryptedPuuid string, apiKey string) (database.SummonerDto, int, error) {
	url := fmt.Sprintf(config.SummonerInfoEndpoint, encryptedPuuid, apiKey)

	if encryptedPuuid == "" {
		return database.SummonerDto{}, http.StatusBadRequest, fmt.Errorf("missing puuid")
	}

	resp, err := http.Get(url)
	if err != nil {
		return database.SummonerDto{}, http.StatusInternalServerError, fmt.Errorf("failed to make HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return database.SummonerDto{}, resp.StatusCode, fmt.Errorf("received non-OK status code: %d", resp.StatusCode)
	}

	var summoner database.SummonerDto
	if err := json.NewDecoder(resp.Body).Decode(&summoner); err != nil {
		log.Println(err)
		return database.SummonerDto{}, http.StatusInternalServerError, fmt.Errorf("failed to decode response: %v", err)
	}

	return summoner, http.StatusOK, nil
}
