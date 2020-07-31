package cmd

import (
	"github.com/vsivarajah/RiotStatistics/pkg/deps"
	"github.com/vsivarajah/RiotStatistics/pkg/setup"
)

func Start() (string, error) {

	dep, message, err := deps.New()
	if err != nil {
		return message, err
	}

	app := setup.New(dep)
	if err := app.Router.Run(":8085"); err != nil {
		return "", err
	}
	return "", nil
}
