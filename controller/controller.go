package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/vsivarajah/RiotStatistics/summoner"
)

var apiKey = "RGAPI-404efd6b-e0a6-4194-bd44-7feb43a7efc0"

func GetSummoner(c *gin.Context) {
	summonerName := c.Param("name")
	summoner.GetSummonerDetails(summonerName)

}
