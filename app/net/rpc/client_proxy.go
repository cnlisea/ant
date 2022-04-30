package rpc

import (
	"context"
	"strconv"
	"sync"

	"github.com/cnlisea/ant/app/proxy"
)

type ClientProxy struct {
	discovery proxy.Discovery
	groupName string
	clients   map[string]*Client
	lock      *sync.RWMutex

	closed bool
	mutex  sync.RWMutex
}

func NewClientProxy(discovery proxy.Discovery, groupName string) *ClientProxy {
	return &ClientProxy{
		discovery: discovery,
		groupName: groupName,
		clients:   make(map[string]*Client),
		lock:      new(sync.RWMutex),
	}
}

func (cp *ClientProxy) GetClient(ctx context.Context, name string, serviceName string, selectMode ClientSelectMode) (*Client, error) {
	key := strconv.FormatUint(uint64(selectMode), 10) + ":" + name + ":" + serviceName
	cp.lock.RLock()
	c := cp.clients[key]
	cp.lock.RUnlock()

	if c == nil {
		newClient, err := NewClient(cp.discovery, name, serviceName, ClientWithSelectMode(selectMode), ClientWithGroupName(cp.groupName))
		if err != nil {
			return nil, err
		}

		var used bool
		cp.lock.Lock()
		c = cp.clients[key]
		if c == nil {
			c = newClient
			cp.clients[key] = c
			used = true
		}
		cp.lock.Unlock()

		// close new client for not used
		if !used {
			if err = newClient.Close(); err != nil {
				return nil, err
			}
		}
	}

	return c, nil
}

func (cp *ClientProxy) Close() {
	cp.mutex.Lock()
	if !cp.closed {
		cp.closed = true
		cp.lock.Lock()
		for _, c := range cp.clients {
			c.Close()
		}
		cp.lock.Unlock()
	}
	cp.mutex.Unlock()
}
