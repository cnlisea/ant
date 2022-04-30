package mq

import (
	"context"

	"github.com/cnlisea/ant/app/mq/client/rocket"
	"github.com/cnlisea/ant/app/mq/message"
	"github.com/cnlisea/ant/app/mq/option"
)

type ConsumerMessage = message.Consumer

type Consumer interface {
	Start(ctx context.Context) error
	Close(ctx context.Context) error
	Subscribe(ctx context.Context, topic string, tag string, f func(context.Context, ...*ConsumerMessage) bool) error
	Unsubscribe(ctx context.Context, topic string) error
}

func NewConsumerRocket(ctx context.Context, nameServer []string, namespace string, groupId string, options ...option.ConsumerFunc) (Consumer, error) {
	return rocket.NewConsumer(ctx, nameServer, namespace, groupId, options...)
}
