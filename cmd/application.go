package cmd

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/vsivarajah/RiotStatistics/api"
	"github.com/vsivarajah/RiotStatistics/pkg/setup"
	"github.com/vsivarajah/RiotStatistics/producer"
	"github.com/vsivarajah/RiotStatistics/producer/db"
	"github.com/vsivarajah/RiotStatistics/producer/kafka"
)

func Start() error {

	// setup client
	cc := &http.Client{Timeout: 5 * time.Second}
	client := api.New(cc)
	key := os.Getenv("RIOTAPI_KEY")
	if key == "" {
		return errors.New("api key missing")
	}
	client.APIKey = key

	// setup kafka producer
	sender := os.Getenv("SENDER")
	var prd producer.Sender
	switch sender {
	case "kafka":
		fmt.Println("kafka producer is run")
		prd = kafka.New()
	default:
		prd = db.New("")
	}

	ctx := context.Background()

	if err := prd.Init(ctx, nil); err != nil {
		return err
	}

	app := setup.New(client, prd)
	if err := app.Router.Run(":8085"); err != nil {
		return err
	}
	return nil
}
