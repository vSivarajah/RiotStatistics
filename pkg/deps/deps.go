package deps

import (
	"net/http"

	"github.com/vsivarajah/RiotStatistics/api"
	"github.com/vsivarajah/RiotStatistics/pkg/config"
	repo "github.com/vsivarajah/RiotStatistics/repositories"
	db "github.com/vsivarajah/RiotStatistics/repositories/db"
)

type Dependencies struct {
	Client       *api.Client
	DbRepository repo.DbRepository
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
	/*
		sender, err := kk.New(c)
		if err != nil {
			return nil, "error setting sender", err
		}
	*/

	dbRepository, err := db.NewRedisCache(c)
	if err != nil {
		return nil, "error setting sender", err
	}

	return &Dependencies{
		Client:       client,
		DbRepository: dbRepository,
	}, "", nil

}
