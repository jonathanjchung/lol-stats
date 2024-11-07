package lol

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/k0kubun/pp"
)

// this one probably goes to client.go
func (rc *RiotClient) TestGettingAccountInfo(w http.ResponseWriter, r *http.Request) {
	summonerName := r.URL.Query().Get("summonerName")
	tagLine := r.URL.Query().Get("tag")

	acc := GetAccountInfo(summonerName, tagLine, rc.apiKey)
	matches := GetMatchHistory(acc.Puuid, rc.apiKey)
	matchData := GetMatchDataFromId(matches[1], rc.apiKey)

	//log.Println(acc)
	//log.Println(matches)
	//log.Println(matchData)
	pp.Print(matchData)

	w.WriteHeader(http.StatusOK)
}

func GetAccountInfo(summonerName string, tagLine string, apiKey string) Account {
	url := fmt.Sprintf("https://americas.api.riotgames.com/riot/account/v1/accounts/by-riot-id/%s/%s?api_key=%s", summonerName, tagLine, apiKey)

	if summonerName == "" || tagLine == "" {
		log.Println("No SummonerName or TagLine")
		return Account{}
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return Account{}
	}
	defer resp.Body.Close()

	var acc Account
	if err := json.NewDecoder(resp.Body).Decode(&acc); err != nil {
		log.Println(err)
		return Account{}
	}

	return acc
}

// Not sure if this one is needed yet
func GetSummonerInfo(encryptedPuuid string, apiKey string) SummonerDto {
	url := fmt.Sprintf("https://na1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/%s?api_key=%s", encryptedPuuid, apiKey)

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
