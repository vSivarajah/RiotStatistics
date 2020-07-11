package summoner

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type successResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
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

	url := fmt.Sprintf("https://euw1.api.riotgames.com/lol/summoner/v4/summoners/by-name/%s", summonerName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// handle err
	}
	req.Header.Set("X-Riot-Token", os.Getenv("API_KEY"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	summonerDetails := Summoner{}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err = json.NewDecoder(resp.Body).Decode(&errRes); err == nil {
			log.Println(err)
			return summonerDetails, errors.New(errRes.Message)

		}
		return summonerDetails, fmt.Errorf("unknown error, status code: %d", resp.StatusCode)
	}
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
	url := fmt.Sprintf("https://euw1.api.riotgames.com/lol/league/v4/entries/by-summoner/%s", summoner.Id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// handle err
	}
	req.Header.Set("X-Riot-Token", os.Getenv("API_KEY"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	summonerProfile := SummonerProfileDetails{}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err = json.NewDecoder(resp.Body).Decode(&errRes); err == nil {
			return summonerProfile, errors.New(errRes.Message)

		}
		return summonerProfile, fmt.Errorf("unknown error, status code: %d", resp.StatusCode)

	}

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
