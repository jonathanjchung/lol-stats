package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
)

func (pg *PostgresInfo) InsertSummonerInfoIntoDB(player *Player) {
	query := `INSERT INTO player_info (puuid, summoner_name, tag_line, account, summoner, matches) VALUES ($1, $2, $3, $4::jsonb, $5::jsonb, $6::jsonb)`

	accountJson, err := json.Marshal(player.Account)
	if err != nil {
		fmt.Println(err)
	}

	summonerJson, err := json.Marshal(player.Summoner)
	if err != nil {
		fmt.Println(err)
	}

	matchJson, err := json.Marshal(player.Matches)
	if err != nil {
		fmt.Println(err)
	}

	_, err = pg.database.Exec(query, player.Account.Puuid, player.Account.GameName, player.Account.TagLine, string(accountJson), string(summonerJson), string(matchJson))

	if err != nil {
		log.Println("insert into table failed")
		log.Println(err)
	} else {
		log.Println("insert into table success")
	}
}

func (pg *PostgresInfo) GetSummonerInfoFromDB(summonerName string, tagLine string) *Player {
	player := new(Player)
	var err error
	var matchJson []byte

	query := `SELECT puuid, summoner_name, tag_line, account, summoner, matches FROM player_info WHERE summoner_name = $1 AND tag_line = $2`
	err = pg.database.QueryRow(query, summonerName, tagLine).Scan(
		&player.Puuid,
		&player.SummonerName,
		&player.TagLine,
		&player.Account,
		&player.Summoner,
		&matchJson,
	)

	if err != nil {
		log.Println("player query error")
		log.Println(err)
		return nil
	}

	matches := []MatchDto{}
	err = json.Unmarshal(matchJson, &matches)

	if err != nil {
		log.Println("match history unmarshal error")
		return nil
	}

	player.Matches = matches

	log.Println("selected summoner info")

	return player
}

func (pg *PostgresInfo) SummonerExists(summonerName string, tagLine string) bool {
	query := `SELECT summoner_name, tag_line FROM player_info WHERE summoner_name = $1 AND tag_line = $2`
	err := pg.database.QueryRow(query, summonerName, tagLine).Scan(&summonerName, &tagLine)

	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return false
	}

	log.Println("summoner does exist")

	return true
}
