package api

import (
	"fmt"

	"github.com/google/go-querystring/query"
	"github.com/vsivarajah/RiotStatistics/utils"
)

const (
	kafkaConn = "localhost:9092"
	topic     = "senz"
)

type Match struct {
	*MatchListDto
	*MatchDTO
}

// MatchList struct describes a response to a Match List api call
type MatchListDto struct {
	Matches    []MatchReferenceDto `json:"matches,omitempty"`
	TotalGames int                 `json:"totalGames"`
	StartIndex int                 `json:"startIndex"`
	EndIndex   int                 `json:"endIndex"`
}

type MatchDTO struct {
	GameId                int                     `json:"gameId" bson:"_id" `
	ParticipantIdentities []ParticipantIdentities `json:"participantidentities" bson:"participantidentities"`
	QueueId               int                     `json:"queueId" bson:"queueId"`
	GameType              string                  `json:"gameType" bson:"gameType"`
	GameDuration          int64                   `json:"gameDuration" bson:"gameDuration"`
	Teams                 []TeamStatsDTO          `json:"teams" bson:"teams"`
	PlatformId            string                  `json:"platformId" bson:"platformId"`
	GameCreation          int64                   `json:"gameCreation" bson:"gameCreation"`
	SeasonId              int                     `json:"seasonId" bson:"seasonId"`
	GameVersion           string                  `json:"gameVersion" bson:"gameVersion"`
	MapId                 int                     `json:"mapId" bson:"mapId"`
	GameMode              string                  `json:"gameMode" bson:"gameMode"`
	Participants          []ParticipantDTO        `json:"participants" bson:"participants"`
}

type MatchReferenceDto struct {
	Lane       string `json:"lane"`
	Champion   int    `json:"champion"`
	PlatformId string `json:"platformId"`
	Timestamp  int64  `json:"timestamp"` // Epoch milliseconds
	Region     string `json:"region"`
	GameId     int    `json:"gameId"`
	Queue      int    `json:"queue"`
	Role       string `json:"role"`
	Season     int    `json:"season"`
}

type MatchListOptions struct {
	// The end time to use for fetching games specified as epoch milliseconds.
	EndTime int64 `url:"endTime,omitempty"`

	// The end index to use for fetching games.
	EndIndex int `url:"endIndex,omitempty"`

	// Set of queue IDs for filtering matchlist.
	Queues []QueueType `url:"queue,omitempty"`

	// Set of season IDs for filtering matchlist.
	Seasons []int `url:"season,omitempty"`

	// The begin time to use for fetching games specified as epoch milliseconds.
	BeginTime int64 `url:"beginTime,omitempty"`

	// The begin index to use for fetching games.
	BeginIndex int `url:"beginIndex,omitempty"`

	// Comma-separated list of champion IDs to use for fetching games.
	Champions []int `url:"champion,omitempty"`
}

type MatchListMethod struct {
	client *Client
}

func (m *MatchListMethod) ByAccount(accountId string, platformId string, options *MatchListOptions) (*MatchListDto, *utils.ApplicationError) {
	relPath := "/lol/match/v4/matchlists/by-account/" + accountId

	if options != nil {
		if vals, err := query.Values(options); err != nil {
			return nil, &utils.ApplicationError{
				Message: err.Error(),
			}
		} else {
			relPath += "?" + vals.Encode()
		}
	}
	data := new(MatchListDto)
	if resp, err := m.client.get(platformURLBase, relPath, platformId, data); err != nil {
		return nil, &utils.ApplicationError{
			StatusCode: resp.StatusCode,
		}
	}

	return data, nil
}

func (m *MatchListMethod) MatchDetailsByGameId(gameId int, platformId string) (*MatchDTO, *utils.ApplicationError) {
	relPath := fmt.Sprintf("/lol/match/v4/matches/%d", gameId)
	fmt.Println(relPath)
	data := new(MatchDTO)

	if resp, err := m.client.get(platformURLBase, relPath, platformId, data); err != nil {
		return nil, &utils.ApplicationError{
			StatusCode: resp.StatusCode,
		}
	}

	return data, nil
}
