package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/vsivarajah/RiotStatistics/utils"
)

var (
	ErrAPIKeyRequired    = errors.New("Required APIKey is missing")
	ErrUnknownPlatformId = errors.New("Unknown platform id")
)

// Version 3+ format, e.g. https://na1.api.riotgames.com/lol/summoner/v3/summoners/by-name/{name}
const platformURLBase = "https://{platform}.api.riotgames.com"

// Version <3 format, e.g.

type Client struct {
	client   *http.Client
	APIKey   string
	Summoner *SummonerMethod
}

func NewClient(httpClient *http.Client) *Client {
	c := &Client{
		client: httpClient,
		APIKey: "",
	}

	c.Summoner = &SummonerMethod{client: c}

	return c
}

func (c *Client) get(basePath, relPath, platformId string, decoded interface{}) (*http.Response, *utils.ApplicationError) {
	if len(c.APIKey) == 0 {
		return nil, &utils.ApplicationError{
			Message: ErrAPIKeyRequired.Error(),
		}
	}

	platform := GetPlatform(platformId)
	if platform == nil {
		return nil, &utils.ApplicationError{
			Message: ErrUnknownPlatformId.Error(),
		}
	}

	relURL, err := url.Parse(replaceTokens(relPath, platform))
	if err != nil {
		return nil, &utils.ApplicationError{
			Message: err.Error(),
		}
	}

	baseURL, _ := url.Parse(replaceTokens(basePath, platform))
	combinedURL := baseURL.ResolveReference(relURL)

	// Add the API Key
	q := combinedURL.Query()
	q.Set("api_key", c.APIKey)
	combinedURL.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", combinedURL.String(), nil)
	fmt.Println(req)

	if err != nil {
		return nil, &utils.ApplicationError{
			Message: err.Error(),
		}
	}

	resp, err := c.client.Do(req)

	if err != nil {
		return nil, &utils.ApplicationError{
			Message: err.Error(),
		}
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNotModified {
		errBody := new(bytes.Buffer)
		resp.Write(errBody)
		return resp, &utils.ApplicationError{
			StatusCode: resp.StatusCode,
		}
	}

	if decoded != nil {
		err = json.NewDecoder(resp.Body).Decode(decoded)
	}

	return resp, nil
}

func replaceTokens(s string, platform *Platform) string {
	s = strings.Replace(s, "{region}", platform.RegionId, -1)
	s = strings.Replace(s, "{platform}", platform.Id, -1)
	return s
}
