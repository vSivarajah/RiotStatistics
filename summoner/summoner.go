package summoner

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var apiKey = "###################"

type Summoner struct {
	Id           string `json:"id"`
	SummonerName string `json:"name"`
}

func GetSummoner(summonerName string) Summoner {

	url := fmt.Sprintf("https://euw1.api.riotgames.com/lol/summoner/v4/summoners/by-name/%s?api_key=%s", summonerName, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	summonerDetails := Summoner{}
	err = json.Unmarshal(body, &summonerDetails)
	if err != nil {
		panic(err)
	}

	return summonerDetails
}

func GetSummonerDetails(summonerName string) {
	summoner := GetSummoner(summonerName)
	url := fmt.Sprintf("https://euw1.api.riotgames.com/lol/league/v4/entries/by-summoner/%s?api_key=%s", summoner.Id, apiKey)
	fmt.Println(summoner.Id, summoner.SummonerName)
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
}
