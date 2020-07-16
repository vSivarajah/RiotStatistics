package app

import "github.com/vsivarajah/RiotStatistics/controller"

func MapUrls() {
	router.GET("/summoner/:name", controller.GetSummoner)
	router.GET("/matches/:id", controller.GetMatchesBySummonerId)
	router.GET("/champions", controller.GetChampions)
	router.GET("/by-summoner/:id", controller.GetPositionsBySummoner)
}
