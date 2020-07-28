package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/vsivarajah/RiotStatistics/controller"
)

func MapUrls() *gin.Engine {
	router := gin.Default()
	router.GET("/matches/:name", controller.GetMatchesBySummonerId)
	router.GET("/champions", controller.GetChampions)
	router.GET("/summoner/:name", controller.GetPositionsBySummoner)
	router.GET("match/:id", controller.GetMatchDetailsByGameId)

	return router
}
