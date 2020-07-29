package champions

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/vsivarajah/RiotStatistics/api"
	"io/ioutil"
	"log"
	"net/http"
)

type Api struct{}

func New() Api {
	return Api{}
}

func (a *Api) Get(ctx *gin.Context) {






	plan, _ := ioutil.ReadFile("champions.json")
	data := api.ChampionData{}
	err := json.Unmarshal(plan, &data)
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, data)
}