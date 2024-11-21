package lol

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jonathanjchung/lol-stats/server/config"
)

func GetAccountInfo(summonerName string, tagLine string, apiKey string) AccountDto {
	url := fmt.Sprintf(config.AccountInfoEndpoint, summonerName, tagLine, apiKey)

	if summonerName == "" || tagLine == "" {
		log.Println("No SummonerName or TagLine")
		return AccountDto{}
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return AccountDto{}
	}
	defer resp.Body.Close()

	var acc AccountDto
	if err := json.NewDecoder(resp.Body).Decode(&acc); err != nil {
		log.Println(err)
		return AccountDto{}
	}

	return acc
}

func GetSummonerInfo(encryptedPuuid string, apiKey string) SummonerDto {
	url := fmt.Sprintf(config.SummonerInfoEndpoint, encryptedPuuid, apiKey)

	if encryptedPuuid == "" {
		log.Println("No puuid")
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	var summoner SummonerDto
	if err := json.NewDecoder(resp.Body).Decode(&summoner); err != nil {
		log.Println(err)
		return SummonerDto{}
	}

	return summoner
}
