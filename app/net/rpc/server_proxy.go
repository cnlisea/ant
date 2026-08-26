package rpc

import "context"

type ServerLocalProxy interface {
	HandlerCall(ctx context.Context, method string, args interface{}, reply interface{}) (exist bool, err error)
}
