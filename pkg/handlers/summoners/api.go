package summoners

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vsivarajah/RiotStatistics/api"
	"net/http"
)

type Api struct {
	client *api.Client
}

func New(c *api.Client) Api {
	return Api{client: c}
}

func (a *Api) GetPositionsBySummoner(c *gin.Context) {
	summonerName := c.Param("name")

	//TODO: This should be fetched from a database if exists, otherwise call RIOT API
	summoner, err := a.client.Summoner.ByName(summonerName, "EUW1")

	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	data, restErr := a.client.League.PositionsBySummoner(summoner.Id, "EUW1")

	fmt.Println(data)
	if restErr != nil {
		c.JSON(restErr.StatusCode, restErr)
		return
	}
	c.JSON(http.StatusOK, data)

}
