package controller

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
	"github.com/vsivarajah/RiotStatistics/api"
)

var (
	client *api.Client
)

func GetSummoner(c *gin.Context) {
	summonerName := c.Param("name")
	client = api.NewClient(new(http.Client))
	client.APIKey = os.Getenv("RIOTAPI_KEY")

	summonerProfile, err := client.Summoner.ByName(summonerName, "EUW1")

	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, summonerProfile)

}
func GetMatchesBySummonerId(c *gin.Context) {
	summonerAccountId := c.Param("id")
	client = api.NewClient(new(http.Client))
	client.APIKey = os.Getenv("RIOTAPI_KEY")
	options := new(api.MatchListOptions)
	minusTwoWeeks, err := time.ParseDuration("-336h")
	if err != nil {
		log.Errorf("MatchList.ByAccount could not parse duration")
	}
	options.BeginTime = time.Now().Add(minusTwoWeeks).Unix() * 1000 // seconds to ms
	options.Queues = make([]api.QueueType, 1)
	options.Queues[0] = api.TEAM_BUILDER_RANKED_SOLO
	matches, errorResponse := client.Matches.ByAccount(summonerAccountId, "EUW1", options)
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
	summonerId := c.Param("id")
	client = api.NewClient(new(http.Client))
	client.APIKey = os.Getenv("RIOTAPI_KEY")
	data, restErr := client.League.PositionsBySummoner(summonerId, "EUW1")
	fmt.Println(data)
	if restErr != nil {
		c.JSON(restErr.StatusCode, restErr)
		return
	}
	c.JSON(http.StatusOK, data)

}
