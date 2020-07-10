package summoner

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var ApiKey = ""

type Summoner struct {
	Id           string `json:"id"`
	SummonerName string `json:"name"`
	AccountId    string `json:"accountid"`
}
type MiniSeries struct {
	Target   int    `json:"target"`
	Wins     int    `json:"wins"`
	Losses   int    `json:"losses"`
	Progress string `json:"progress"`
}
type SummonerProfile struct {
	LeagueId     string `json:"leagueid"`
	QueueType    string `json:"queueType"`
	Tier         string `json:"tier"`
	Rank         string `json:"rank"`
	SummonerId   string `json:"summonerid"`
	SummonerName string `json:"summonername"`
	LeaguePoints int    `json:"leaguepoints"`
	Wins         int    `json:"wins"`
	Losses       int    `json:"losses"`
	Veteran      bool   `json:"veteran"`
	Inactive     bool   `json:"inactive"`
	FreshBlood   bool   `json:"freshblood"`
	HotStreak    bool   `json:"hotstreak"`
	MiniSeries   MiniSeries
}

type SummonerProfileDetails []*SummonerProfile

func GetSummoner(summonerName string) (Summoner, error) {

	url := fmt.Sprintf("https://euw1.api.riotgames.com/lol/summoner/v4/summoners/by-name/%s?api_key=%s", summonerName, ApiKey)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	summonerDetails := Summoner{}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &summonerDetails)
	if err != nil {
		panic(err)
	}

	return summonerDetails, nil
}

func GetSummonerDetails(summonerName string) (SummonerProfileDetails, error) {
	summoner, _ := GetSummoner(summonerName)
	url := fmt.Sprintf("https://euw1.api.riotgames.com/lol/league/v4/entries/by-summoner/%s?api_key=%s", summoner.Id, ApiKey)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	defer resp.Body.Close()
	summonerProfile := SummonerProfileDetails{}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &summonerProfile)
	if err != nil {
		log.Fatal(err)
	}
	return summonerProfile, nil
}
