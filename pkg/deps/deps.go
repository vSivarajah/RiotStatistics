package deps

import (
	"github.com/vsivarajah/RiotStatistics/api"
	"github.com/vsivarajah/RiotStatistics/pkg/config"
	"github.com/vsivarajah/RiotStatistics/producer"
	kk "github.com/vsivarajah/RiotStatistics/producer/kafka"

	"net/http"
)

type Dependencies struct {
	Client *api.Client
	Sender producer.Sender
}

func New() (*Dependencies, string, error) {

	// setup configurations
	c, field, err := config.New()
	if err != nil {
		return nil, field, err
	}

	// here we can setup configurations based on environment
	// for example if ENV is test add mock producer and mock clients
	// if it is local then set up localhost and so on

	cc := &http.Client{Timeout: c.Riot.Timeout}
	client := api.New(cc)
	client.APIKey = c.Riot.ApiKey

	sender, err := kk.New(c)
	if err != nil {
		return nil, "error setting sender", err
	}

	return &Dependencies{
		Client: client,
		Sender: sender,
	}, "", nil

}
