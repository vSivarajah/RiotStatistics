package api

import (
	"github.com/vsivarajah/RiotStatistics/utils"
)

type MiniSeries struct {
	Target   int    `json:"target"`
	Wins     int    `json:"wins"`
	Losses   int    `json:"losses"`
	Progress string `json:"progress"`
}
type SummonerProfile struct {
	LeagueId     string     `json:"leagueId"`
	QueueType    string     `json:"queueType"`
	Tier         string     `json:"tier"`
	Rank         string     `json:"rank"`
	SummonerId   string     `json:"summonerId"`
	SummonerName string     `json:"summonerName"`
	LeaguePoints int        `json:"leaguePoints"`
	Wins         int        `json:"wins"`
	Losses       int        `json:"losses"`
	Veteran      bool       `json:"veteran"`
	Inactive     bool       `json:"inactive"`
	FreshBlood   bool       `json:"freshBlood"`
	HotStreak    bool       `json:"hotStreak"`
	MiniSeries   MiniSeries `json:"miniSeries"`
}
type SummonerProfileDetails []*SummonerProfile

type LeagueMethod struct {
	client *Client
}

func (m *LeagueMethod) PositionsBySummoner(summonerId string, platformId string) (*SummonerProfileDetails, *utils.ApplicationError) {
	relPath := "/lol/league/v4/entries/by-summoner/" + summonerId
	data := new(SummonerProfileDetails)
	if resp, err := m.client.get(platformURLBase, relPath, platformId, data); err != nil {
		return nil, &utils.ApplicationError{
			StatusCode: resp.StatusCode,
		}
	}
	return data, nil
}


