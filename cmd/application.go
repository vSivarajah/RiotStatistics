package cmd

import (
	"errors"
	"github.com/vsivarajah/RiotStatistics/api"
	"github.com/vsivarajah/RiotStatistics/pkg/setup"
	"net/http"
	"os"
	"time"
)

func Start() error {

	cc := &http.Client{
		Timeout: 5 * time.Second,
	}

	client := api.New(cc)

	key := os.Getenv("RIOTAPI_KEY")
	if key == "" {
		return errors.New("api key missing")
	}

	client.APIKey = key

	app := setup.New(client)
	if err := app.Router.Run(":8081"); err != nil {
		return err
	}
	return nil
}
