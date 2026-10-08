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

	ClientInvoke(ctx context.Context, name string, serviceName string, method string, args interface{}, reply interface{}) error
	ClientOneWay(ctx context.Context, name string, serviceName string, method string, args interface{}) error
	ClientBroadcast(ctx context.Context, name string, serviceName string, method string, args interface{}, reply interface{}) error
	ClientCallback(ctx context.Context, name string, serviceName string, method string, args interface{}, reply interface{}, f func(reply interface{}, err error))

	ClientHashWithKey(ctx context.Context, val string) context.Context
	ClientHashInvoke(ctx context.Context, name string, serviceName string, method string, args interface{}, reply interface{}) error
	ClientHashOneWay(ctx context.Context, name string, serviceName string, method string, args interface{}) error
	ClientHashBroadcast(ctx context.Context, name string, serviceName string, method string, args interface{}, reply interface{}) error
	ClientHashCallback(ctx context.Context, name string, serviceName string, method string, args interface{}, reply interface{}, f func(reply interface{}, err error))

	ClientSoftStateWithKey(ctx context.Context, val string) context.Context
	ClientSoftStateInvoke(ctx context.Context, name string, serviceName string, method string, args interface{}, reply interface{}) error
	ClientSoftStateOneWay(ctx context.Context, name string, serviceName string, method string, args interface{}) error
	ClientSoftStateBroadcast(ctx context.Context, name string, serviceName string, method string, args interface{}, reply interface{}) error
	ClientSoftStateCallback(ctx context.Context, name string, serviceName string, method string, args interface{}, reply interface{}, f func(reply interface{}, err error))
}
