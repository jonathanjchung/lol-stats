package database

import (
	"encoding/json"
	"fmt"
)

type Player struct {
	Puuid        string      `db:"puuid"`
	SummonerName string      `db:"summoner_name"`
	TagLine      string      `db:"tag_line"`
	Account      AccountDto  `db:"account"`
	Summoner     SummonerDto `db:"summoner"`
	Matches      []MatchDto  `db:"matches"`
}

func (a *AccountDto) Scan(src any) error {
	data, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("scan: expected []byte, got %T", src)
	}

	return json.Unmarshal(data, &a)
}

func (s *SummonerDto) Scan(src any) error {
	data, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("scan: expected []byte, got %T", src)
	}

	return json.Unmarshal(data, &s)
}
