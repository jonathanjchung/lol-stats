package lol

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jonathanjchung/lol-stats/server/config"
)

func GetAccountInfo(summonerName string, tagLine string, apiKey string) (AccountDto, int, error) {
	url := fmt.Sprintf(config.AccountInfoEndpoint, summonerName, tagLine, apiKey)

	if summonerName == "" || tagLine == "" {
		return AccountDto{}, http.StatusBadRequest, fmt.Errorf("missing summonerName or tagLine")
	}

	resp, err := http.Get(url)
	if err != nil {
		return AccountDto{}, http.StatusInternalServerError, fmt.Errorf("failed to make HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return AccountDto{}, resp.StatusCode, fmt.Errorf("received non-OK status code: %d", resp.StatusCode)
	}

	var acc AccountDto
	if err := json.NewDecoder(resp.Body).Decode(&acc); err != nil {
		return AccountDto{}, http.StatusInternalServerError, fmt.Errorf("failed to decode response: %v", err)
	}

	return acc, http.StatusOK, nil
}

func GetSummonerInfo(encryptedPuuid string, apiKey string) (SummonerDto, int, error) {
	url := fmt.Sprintf(config.SummonerInfoEndpoint, encryptedPuuid, apiKey)

	if encryptedPuuid == "" {
		return SummonerDto{}, http.StatusBadRequest, fmt.Errorf("missing puuid")
	}

	resp, err := http.Get(url)
	if err != nil {
		return SummonerDto{}, http.StatusInternalServerError, fmt.Errorf("failed to make HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return SummonerDto{}, resp.StatusCode, fmt.Errorf("received non-OK status code: %d", resp.StatusCode)
	}

	var summoner SummonerDto
	if err := json.NewDecoder(resp.Body).Decode(&summoner); err != nil {
		log.Println(err)
		return SummonerDto{}, http.StatusInternalServerError, fmt.Errorf("failed to decode response: %v", err)
	}

	return summoner, http.StatusOK, nil
}
