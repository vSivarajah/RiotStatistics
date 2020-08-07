package matchers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/vsivarajah/RiotStatistics/pkg/deps"

	"github.com/gin-gonic/gin"
	"github.com/vsivarajah/RiotStatistics/api"
	repo "github.com/vsivarajah/RiotStatistics/repositories"
)

type Api struct {
	client       *api.Client
	dbRepository repo.DbRepository
}

func New(d *deps.Dependencies) Api {
	return Api{client: d.Client, dbRepository: d.DbRepository}
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
	for _, v := range matches.Matches {

		matchDetail, err := a.client.Matches.MatchDetailsByGameId(v.GameId, "EUW1")
		if err != nil {
			return
		}
		fmt.Println(matchDetail)
	}
	c.JSON(http.StatusOK, matches)

}

func (a *Api) GetMatchDetailsByGameId(c *gin.Context) {

	gameId := c.Param("id")
	gameIdInt, _ := strconv.Atoi(gameId)

	matchDetail := a.dbRepository.Get(c.Request.Context(), gameIdInt)
	if matchDetail == nil {
		log.Println("No data in cache, fetching from 3rd party api")
		matchDetail, err := a.client.Matches.MatchDetailsByGameId(gameIdInt, "EUW1")

		match := &api.Match{MatchDTO: matchDetail}

		if err != nil {
			c.JSON(err.StatusCode, err)
			return
		}
		log.Println("Sending data to cache")
		if err := a.dbRepository.Send(c.Request.Context(), match); err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, match)

	} else {
		log.Println("Data exists in the cache, returning data from cache")
		c.JSON(http.StatusOK, matchDetail.MatchDTO)
	}
}
