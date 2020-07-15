package api

import (
	"fmt"

	"github.com/vsivarajah/RiotStatistics/utils"
)

type SummonerDTO struct {
	AccountId     string `json:"accountId"`
	ProfileIconId int    `json:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate"`
	Name          string `json:"name"`
	Id            string `json:"id"`
	Puuid         string `json:"puuid"`
	SummonerLevel int64  `json:"summonerLevel"`
}

type SummonerMethod struct {
	client *Client
}

func (m *SummonerMethod) ByName(name, platformId string) (*SummonerDTO, *utils.ApplicationError) {
	relPath := "/lol/summoner/v4/summoners/by-name/" + name
	data := new(SummonerDTO)

	if resp, err := m.client.get(platformURLBase, relPath, platformId, data); err != nil {

		return nil, &utils.ApplicationError{StatusCode: resp.StatusCode}
	}
	fmt.Println(data.AccountId)
	return data, nil
}
