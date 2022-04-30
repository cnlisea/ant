package proxy

import "context"

type MQ interface {
	SendMsg(ctx context.Context, name string, topic string, msg []byte) error
	SendMsgSharding(ctx context.Context, name string, key string, topic string, msg []byte) error
}
