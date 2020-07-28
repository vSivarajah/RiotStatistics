package producer

import "context"

type Sender interface {
	Init(ctx context.Context, config interface{}) error
	Send(ctx context.Context, message interface{}) error
}
