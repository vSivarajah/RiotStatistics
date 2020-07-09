package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vsivarajah/RiotStatistics/champions"
	"github.com/vsivarajah/RiotStatistics/matches"
	"github.com/vsivarajah/RiotStatistics/summoner"
)

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

func GetChampions(c *gin.Context) {
	champions := champions.GetChampions()
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, champions)
}
