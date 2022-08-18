package runtime

import (
	"context"
	"io"
)

type MessageBrokerProducer interface {
	Write(ctx context.Context, msg []byte) error
}

type MessageHandler func(ctx context.Context, msg []byte) error

type MessageBrokerConsumer interface {
	Handle(ctx context.Context, handler MessageHandler) (io.Closer, error)
}
