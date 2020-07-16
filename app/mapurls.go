package app

import "github.com/vsivarajah/RiotStatistics/controller"

func MapUrls() {
	router.GET("/matches/:name", controller.GetMatchesBySummonerId)
	router.GET("/champions", controller.GetChampions)
	router.GET("/summoner/:name", controller.GetPositionsBySummoner)
}
