package producer

import "context"

type Sender interface {
	Send(ctx context.Context, message interface{}) error
}
