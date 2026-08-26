package rpc

import (
	"context"
	"errors"
	"sync"

	"github.com/cnlisea/ant/app/net/rpc/selector"
	"github.com/cnlisea/ant/app/proxy"
	rpcNacosClient "github.com/rpcxio/rpcx-nacos/client"
	"github.com/smallnest/rpcx/client"
)

type Client struct {
	ServerLocalProxy ServerLocalProxy

	client client.XClient
	closed bool
	mutex  sync.RWMutex
}

func NewClient(discovery proxy.Discovery, name string, serviceName string, options ...ClientOptionFunc) (*Client, error) {
	if discovery == nil {
		return nil, errors.New("discovery invalid")
	}

	var op ClientOption
	for _, f := range options {
		f(&op)
	}

	var (
		failMode         client.FailMode
		selectMode       client.SelectMode
		selectorInstance client.Selector
	)
	switch op.SelectMode {
	case ClientSelectModeHash:
		failMode = client.Failtry
		selectMode = client.SelectByUser
		selectorInstance = new(selector.Hash)
	case ClientSelectModeSoftState:
		failMode = client.Failtry
		selectMode = client.SelectByUser
		selectorInstance = new(selector.SoftState)
	default:
		failMode = client.Failover
		selectMode = client.RoundRobin
	}

	xClient := client.NewXClient(serviceName,
		failMode,
		selectMode,
		rpcNacosClient.NewNacosDiscoveryWithClient(name, discovery.ClusterName(), op.GroupName, discovery.Instance()),
		client.DefaultOption)
	if selectorInstance != nil {
		xClient.SetSelector(selectorInstance)
	}

	return &Client{
		client:           xClient,
		ServerLocalProxy: op.ServerLocalProxy,
	}, nil
}

func (c *Client) WithHashKey(ctx context.Context, val string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, selector.CtxHashTag, val)
}

func (c *Client) WithSoftStateKey(ctx context.Context, val string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, selector.CtxSoftStateTag, val)
}

func (c *Client) Invoke(ctx context.Context, method string, args interface{}, reply interface{}) error {
	if c.ServerLocalProxy != nil {
		exist, err := c.ServerLocalProxy.HandlerCall(ctx, method, args, reply)
		if exist {
			return err
		}
	}
	return c.client.Call(ctx, method, args, reply)
}

func (c *Client) OneWay(ctx context.Context, method string, args interface{}) error {
	if c.ServerLocalProxy != nil {
		exist, err := c.ServerLocalProxy.HandlerCall(ctx, method, args, nil)
		if exist {
			return err
		}
	}
	return c.client.Call(ctx, method, args, nil)
}

func (c *Client) Broadcast(ctx context.Context, method string, args interface{}, reply interface{}) error {
	return c.client.Broadcast(ctx, method, args, reply)
}

func (c *Client) Callback(ctx context.Context, method string, args interface{}, reply interface{}, f func(reply interface{}, err error)) {
	go func() {
		if c.ServerLocalProxy != nil {
			exist, err := c.ServerLocalProxy.HandlerCall(ctx, method, args, reply)
			if exist {
				f(reply, err)
				return
			}
		}
		f(reply, c.client.Call(ctx, method, args, reply))
	}()
}

func (c *Client) Close() error {
	var err error
	c.mutex.Lock()
	if !c.closed {
		c.closed = true
		err = c.client.Close()
	}
	c.mutex.Unlock()
	return err
}
