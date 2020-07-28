package cmd

import (
	"context"
	"errors"
	"github.com/vsivarajah/RiotStatistics/api"
	"github.com/vsivarajah/RiotStatistics/pkg/setup"
	"github.com/vsivarajah/RiotStatistics/producer"
	"github.com/vsivarajah/RiotStatistics/producer/db"
	"github.com/vsivarajah/RiotStatistics/producer/kafka"
	"net/http"
	"os"
	"time"
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
		prd = kafka.New()
	default:
		prd = db.New("")
	}

	ctx := context.Background()

	if err := prd.Init(ctx, nil); err != nil {
		return err
	}

	//ctxChild, cancel := context.WithCancel(ctx)
	//defer cancel()

	//timeout,cancel := context.WithTimeout(ctx, time.Second * 5)
	//defer cancel()

	// set timeout as context

	// some error happens

	// setup kafka  consumer
	app := setup.New(client, prd)
	if err := app.Router.Run(":8081"); err != nil {
		return err
	}
	return nil
}

// passing values from hendler --> downstream

// client ---> server
