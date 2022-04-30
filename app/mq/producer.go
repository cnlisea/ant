package mq

import (
	"context"

	"github.com/cnlisea/ant/app/mq/client/rocket"
	"github.com/cnlisea/ant/app/mq/message"
	"github.com/cnlisea/ant/app/mq/option"
)

type ProducerMessage = message.Producer

type Producer interface {
	Start(ctx context.Context) error
	Close(ctx context.Context) error
	SendMsg(ctx context.Context, topic string, msg *ProducerMessage) error
}

func NewProducerRocket(ctx context.Context, nameServer []string, namespace string, groupId string, options ...option.ProducerFunc) (Producer, error) {
	return rocket.NewProducer(ctx, nameServer, namespace, groupId, options...)
}
