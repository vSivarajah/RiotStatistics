package mock

import (
	"context"
	"github.com/vsivarajah/RiotStatistics/producer"
)

type mock struct {
	items []interface{}
}

func New() producer.Sender {
	return &mock{}
}

func (m *mock) Send(ctx context.Context, message interface{}) error {
	m.items = append(m.items, message)
	return nil
}
