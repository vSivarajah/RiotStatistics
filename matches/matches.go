package matches

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/vsivarajah/RiotStatistics/summoner"
)

var ApiKey string = "RGAPI-730fd948-b874-4070-ba8b-82f5bbad6cf6"

type Matches struct {
	Totalgames int     `json:"totalgames"`
	Startindex int     `json:"startindex"`
	Endindex   int     `json:"endindex"`
	Match      []Match `json:"matches"`
}

type Match struct {
	PlatformId string `json:"platformid"`
	GameId     int64  `json:"gameid"`
	Champion   int    `json:"champion"`
	Queue      int    `json:"queue"`
	Season     int    `json:"season"`
	Timestamp  int64  `json:"timestamp"`
	Role       string `json:"role"`
	Lane       string `json:"lane"`
}

type MatchList []*Match

func GetMatches(summonerName string) Matches {
	summoner := summoner.GetSummoner(summonerName)
	url := fmt.Sprintf("https://euw1.api.riotgames.com/lol/match/v4/matchlists/by-account/%s?api_key=%s", summoner.AccountId, ApiKey)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	matches := Matches{}
	err = json.Unmarshal(body, &matches)
	if err != nil {
		log.Fatal(err)
	}
	return matches
}
