package config

const (
	AccountInfoEndpoint  = "https://americas.api.riotgames.com/riot/account/v1/accounts/by-riot-id/%s/%s?api_key=%s"
	SummonerInfoEndpoint = "https://na1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/%s?api_key=%s"
	MatchHistoryEndpoint = "https://americas.api.riotgames.com/lol/match/v5/matches/by-puuid/%s/ids?start=0&count=20&api_key=%s"
	MatchDataEndpoint    = "https://americas.api.riotgames.com/lol/match/v5/matches/%s?api_key=%s"
)
