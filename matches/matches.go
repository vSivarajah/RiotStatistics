package matches

var ApiKey = ""

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

/*
func GetMatches(summonerName string) (Matches, error) {
	summoner, _ := summoner.GetSummoner(summonerName)
	url := fmt.Sprintf("https://euw1.api.riotgames.com/lol/match/v4/matchlists/by-account/%s", summoner.AccountId)
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
	matches := Matches{}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		var errRes utils.ErrorResponse
		if err = json.NewDecoder(resp.Body).Decode(&errRes); err == nil {
			return matches, errors.New(errRes.Message)

		}
		return matches, fmt.Errorf("unknown error, status code: %d", resp.StatusCode)

	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	err = json.Unmarshal(body, &matches)
	if err != nil {
		log.Fatal(err)
	}
	return matches, nil
}

*/
