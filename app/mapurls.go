package app

import "github.com/vsivarajah/RiotStatistics/controller"

func MapUrls() {
	router.GET("/summoner/:name", controller.GetSummoner)
	router.GET("/summoner/:name/matches", controller.GetSummonerMatches)
}
