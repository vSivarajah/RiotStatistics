package controller

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vsivarajah/RiotStatistics/api"
)

var (
	client *api.Client
)

func GetMatchesBySummonerId(c *gin.Context) {
	summonerName := c.Param("name")
	client = api.NewClient(new(http.Client))
	client.APIKey = os.Getenv("RIOTAPI_KEY")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if client.APIKey == "" {
		c.JSON(http.StatusBadRequest, "API KEY not provided")
		return
	}

	//TODO: This should be fetched from a database if exists, otherwise call RIOT API
	summoner, err := client.Summoner.ByName(summonerName, "EUW1")
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	options := new(api.MatchListOptions)
	minusTwoWeeks, _ := time.ParseDuration("-336h")

	options.BeginTime = time.Now().Add(minusTwoWeeks).Unix() * 1000 // seconds to ms
	options.Queues = make([]api.QueueType, 1)
	options.Queues[0] = api.TEAM_BUILDER_RANKED_SOLO
	matches, errorResponse := client.Matches.ByAccount(summoner.AccountId, "EUW1", options)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, matches)

}

func GetChampions(c *gin.Context) {
	champions := api.GetChampions()
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, champions)
}

func GetPositionsBySummoner(c *gin.Context) {
	summonerName := c.Param("name")
	client = api.NewClient(new(http.Client))
	client.APIKey = os.Getenv("RIOTAPI_KEY")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if client.APIKey == "" {
		c.JSON(http.StatusBadRequest, "API KEY not provided")
		return
	}

	//TODO: This should be fetched from a database if exists, otherwise call RIOT API
	summoner, err := client.Summoner.ByName(summonerName, "EUW1")
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	data, restErr := client.League.PositionsBySummoner(summoner.Id, "EUW1")
	fmt.Println(data)
	if restErr != nil {
		c.JSON(restErr.StatusCode, restErr)
		return
	}
	c.JSON(http.StatusOK, data)

}

func GetMatchDetailsByGameId(c *gin.Context) {
	gameId := c.Param("id")
	gameIdInt, _ := strconv.Atoi(gameId)
	client = api.NewClient(new(http.Client))
	client.APIKey = os.Getenv("RIOTAPI_KEY")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if client.APIKey == "" {
		c.JSON(http.StatusBadRequest, "API KEY not provided")
		return
	}

	matchDetail, err := client.Matches.MatchDetailsByGameId(gameIdInt, "EUW1")
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, matchDetail)
}
