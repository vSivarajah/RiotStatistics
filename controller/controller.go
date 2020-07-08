package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vsivarajah/RiotStatistics/matches"
	"github.com/vsivarajah/RiotStatistics/summoner"
)

var apiKey = "RGAPI-404efd6b-e0a6-4194-bd44-7feb43a7efc0"

func GetSummoner(c *gin.Context) {
	summonerName := c.Param("name")
	summonerProfile := summoner.GetSummonerDetails(summonerName)
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, summonerProfile)

}

func GetSummonerMatches(c *gin.Context) {
	summonerName := c.Param("name")
	summonerMatches := matches.GetMatches(summonerName)
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, summonerMatches)

}
