package api

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPositionsBySummonerNoError(t *testing.T) {

	//Initialization
	summonerId := "Koa6GIXPMvhQ6BW402ph5mM02FXSVBEKOWW3p2bbZNbhmYQ"
	client := NewClient(new(http.Client))
	client.APIKey = os.Getenv("RIOTAPI_KEY")

	//Execution
	data, restErr := client.League.PositionsBySummoner(summonerId, "EUW1")

	//Validation

	if client.APIKey == "" {
		t.Errorf("API Key was not specified")
	}

	assert.Nil(t, restErr)
	assert.NotNil(t, data)
}

func TestPositionsBySummonerError(t *testing.T) {

	//Initialization
	summonerId := "Koa6"
	client := NewClient(new(http.Client))
	client.APIKey = os.Getenv("RIOTAPI_KEY")

	//Execution
	data, restErr := client.League.PositionsBySummoner(summonerId, "EUW1")

	//Validation

	if client.APIKey == "" {
		t.Fatalf("API Key was not specified")
	}

	assert.Nil(t, data)
	assert.NotNil(t, restErr)
}
