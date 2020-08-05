package producer

import (
	"context"

	"github.com/vsivarajah/RiotStatistics/api"
)

type DbRepository interface {
	Send(ctx context.Context, match *api.Match) error
	Get(ctx context.Context, key int) *api.Match
}
