package proxy

import (
	"context"
	"github.com/cnlisea/ant/app/mq/message"
)

type MQConsumerSubscribe struct {
	Topic   string
	Tag     string
	Handler func(context.Context, ...*message.Consumer) bool
}

type MQ interface {
	SendMsg(ctx context.Context, name string, topic string, msg []byte) error
	SendMsgSharding(ctx context.Context, name string, key string, topic string, msg []byte) error
}
