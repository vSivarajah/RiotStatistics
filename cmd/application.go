package cmd

import (
	"fmt"
	"github.com/vsivarajah/RiotStatistics/pkg/config"
	"net/http"
	"os"
	"time"

	"github.com/vsivarajah/RiotStatistics/api"
	"github.com/vsivarajah/RiotStatistics/pkg/setup"
	"github.com/vsivarajah/RiotStatistics/producer"
	"github.com/vsivarajah/RiotStatistics/producer/db"
	"github.com/vsivarajah/RiotStatistics/producer/kafka"
)

func Start() (string, error) {

	// setup configurations
	conf, field, err := config.New()
	if err != nil {
		return field, err
	}

	// setup client
	cc := &http.Client{Timeout: 5 * time.Second}
	client := api.New(cc)
	client.APIKey = conf.Riot.ApiKey

	// setup kafka producer
	sender := os.Getenv("SENDER")
	var prd producer.Sender
	switch sender {
	case "kafka":
		fmt.Println("kafka producer is run")
		prd, err = kafka.New(conf)
		if err != nil {
			return "kafka config failed", err
		}
	default:
		prd = db.New("")
	}

	app := setup.New(client, prd)
	if err := app.Router.Run(":8085"); err != nil {
		return "", err
	}
	return "", nil
}
