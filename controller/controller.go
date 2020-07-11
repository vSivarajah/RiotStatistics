package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vsivarajah/RiotStatistics/champions"
	"github.com/vsivarajah/RiotStatistics/matches"
	"github.com/vsivarajah/RiotStatistics/rest_errors"
	"github.com/vsivarajah/RiotStatistics/summoner"
)

func GetSummoner(c *gin.Context) {
	summonerName := c.Param("name")
	summonerProfile, err := summoner.GetSummonerDetails(summonerName)
	fmt.Println(c.Request)
	if err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	c.JSON(http.StatusOK, summonerProfile)

}

func GetSummonerMatches(c *gin.Context) {
	summonerName := c.Param("name")
	summonerMatches, err := matches.GetMatches(summonerName)
	if err != nil {
		restErr := rest_errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}
	c.JSON(http.StatusOK, summonerMatches)

}

func GetChampions(c *gin.Context) {
	champions := champions.GetChampions()
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, champions)
}
