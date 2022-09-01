package runtime

import (
	"context"
)

type MessageBrokerProducer interface {
	Write(ctx context.Context, msg []byte) error
}

type MessageHandler func(ctx context.Context, msg []byte) error

type MessageBrokerConsumer interface {
	Handle(ctx context.Context, handler MessageHandler) error
}
