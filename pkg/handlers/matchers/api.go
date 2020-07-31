package matchers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/vsivarajah/RiotStatistics/pkg/deps"

	"github.com/gin-gonic/gin"
	"github.com/vsivarajah/RiotStatistics/api"
	"github.com/vsivarajah/RiotStatistics/producer"
)

type Api struct {
	client *api.Client
	sender producer.Sender
}

func New(d *deps.Dependencies) Api {
	return Api{client: d.Client, sender: d.Sender}
}

func (a *Api) GetMatchesBySummonerId(c *gin.Context) {
	summonerName := c.Param("name")

	//TODO: This should be fetched from a database if exists, otherwise call RIOT API
	summoner, err := a.client.Summoner.ByName(summonerName, "EUW1")
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	options := new(api.MatchListOptions)
	minusTwoWeeks, _ := time.ParseDuration("-336h")

	options.BeginTime = time.Now().Add(minusTwoWeeks).Unix() * 1000 // seconds to ms
	options.Queues = make([]api.QueueType, 1)
	options.Queues[0] = api.TEAM_BUILDER_RANKED_SOLO
	matches, errorResponse := a.client.Matches.ByAccount(summoner.AccountId, "EUW1", options)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, matches)

}

func (a *Api) GetMatchDetailsByGameId(c *gin.Context) {

	gameId := c.Param("id")
	gameIdInt, _ := strconv.Atoi(gameId)

	matchDetail, err := a.client.Matches.MatchDetailsByGameId(gameIdInt, "EUW1")
	fmt.Println("this method gets called, GetMatchDetails")
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	/*
		if err := a.sender.Send(c.Request.Context(), matchDetail); err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
	*/
	c.JSON(http.StatusOK, matchDetail)
}
