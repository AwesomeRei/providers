package pubsub

import "context"

type EventStreamer interface {
	SendEvent(ctx context.Context,topic string,msg MessageBox) error
}
