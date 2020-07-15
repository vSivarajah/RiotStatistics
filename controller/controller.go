package controller

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vsivarajah/RiotStatistics/api"
)

var (
	client *api.Client
)

func GetSummoner(c *gin.Context) {
	summonerName := c.Param("name")
	client = api.NewClient(new(http.Client))
	client.APIKey = os.Getenv("RIOTAPI_KEY")

	summonerProfile, err := client.Summoner.ByName(summonerName, "EUW1")

	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, summonerProfile)

}

/*
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

*/
