package setup

import (
	"github.com/gin-gonic/gin"
	"github.com/vsivarajah/RiotStatistics/api"
	"github.com/vsivarajah/RiotStatistics/pkg/handlers/champions"
	"github.com/vsivarajah/RiotStatistics/pkg/handlers/matchers"
	"github.com/vsivarajah/RiotStatistics/pkg/handlers/summoners"
	"github.com/vsivarajah/RiotStatistics/pkg/middlewares"
)

type App struct {
	// add all handlers and other configs
	Router   *gin.Engine
	apiChamp champions.Api
	apiMatch matchers.Api
	apiSumnr summoners.Api
}

func New(c *api.Client) *App {
	app := new(App)

	// setup all apis
	app.apiChamp = champions.New()
	app.apiMatch = matchers.New(c)
	app.apiSumnr = summoners.New(c)

	// setup routes
	app.Router = gin.Default()
	root := app.Router.Group("/")
	root.Use(middlewares.CustomHeaders())
	{

		root.GET("/matches/:name", app.apiMatch.GetMatchesBySummonerId)
		root.GET("/champions", app.apiChamp.Get)
		root.GET("/summoner/:name", app.apiSumnr.GetPositionsBySummoner)
		root.GET("match/:id", app.apiMatch.GetMatchDetailsByGameId)

	}

	return app
}
