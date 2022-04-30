package proxy

import (
	"context"
)

type RpcClient interface {
	WithHashKey(ctx context.Context, val string) context.Context
	WithSoftStateKey(ctx context.Context, val string) context.Context
	Invoke(ctx context.Context, method string, args interface{}, reply interface{}) error
	OneWay(ctx context.Context, method string, args interface{}) error
	Broadcast(ctx context.Context, method string, args interface{}, reply interface{}) error
	Callback(ctx context.Context, method string, args interface{}, reply interface{}, f func(reply interface{}, err error))
}

type RpcClientPool interface {
	GetClient(name string, serviceName string) (RpcClient, error)
	GetClientHash(name string, serviceName string) (RpcClient, error)
	GetClientSoftState(name string, serviceName string) (RpcClient, error)
}
