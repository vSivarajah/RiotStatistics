package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/vsivarajah/RiotStatistics/controller"
)

func MapUrls() *gin.Engine {
	router := gin.Default()

	root := router.Group("/")
	root.Use(CustomHeaders())
	{

		root.GET("/matches/:name", controller.GetMatchesBySummonerId)
		root.GET("/champions", controller.GetChampions)
		root.GET("/summoner/:name", controller.GetPositionsBySummoner)
		root.GET("match/:id", controller.GetMatchDetailsByGameId)

	}

	return router
}
