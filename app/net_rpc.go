package app

import (
	"sync"

	netRpc "github.com/cnlisea/ant/app/net/rpc"
	"github.com/cnlisea/ant/logs"
)

func (a *App) NetRpcRegister(name string, ip string, port uint16, groupName string, discoverySoftState *bool, service interface{}) error {
	if a.rpcServers == nil {
		a.rpcServers = make(map[string]*netRpc.Server, 1)
	}

	s := a.rpcServers[name]
	if s == nil {
		s = netRpc.NewServer(ip, port)
		a.rpcServers[name] = s

		if groupName != "" {
			if a.rpcServerGroupName == nil {
				a.rpcServerGroupName = make(map[string]string, 1)
			}
			a.rpcServerGroupName[name] = groupName
		}
	}

	if discoverySoftState != nil {
		a._DiscoverySoftStateAdd(_DiscoverySoftStateRpcNamePrefix, name, *discoverySoftState)
	}
	return s.SetHandler(service)
}

type NetRpcClientOption struct {
	SelectMode netRpc.ClientSelectMode
}

type NetRpcClientOptionFunc func(*NetRpcClientOption)

func (a *App) NetRpcClientWithHashSelector() NetRpcClientOptionFunc {
	return func(op *NetRpcClientOption) {
		op.SelectMode = netRpc.ClientSelectModeHash
	}
}

func (a *App) NetRpcClientWithSoftState() NetRpcClientOptionFunc {
	return func(op *NetRpcClientOption) {
		op.SelectMode = netRpc.ClientSelectModeSoftState
	}
}

func (a *App) NetRpcClient(groupName string, name string, serviceName string, options ...NetRpcClientOptionFunc) (*netRpc.Client, error) {
	var op NetRpcClientOption
	for _, f := range options {
		f(&op)
	}
	if a.rpcClientProxy == nil {
		a.rpcClientProxy = netRpc.NewClientProxy(a.ProxyDiscovery(), groupName)
	}
	return a.rpcClientProxy.GetClient(a.Context(), name, serviceName, op.SelectMode)
}

func (a *App) NetRpcAllStart(ws *sync.WaitGroup) bool {
	var (
		name string
		s    *netRpc.Server
	)
	for name, s = range a.rpcServers {
		ws.Add(1)
		go func(name string, server *netRpc.Server) {
			logs.Info("rpc server run",
				logs.String("name", name),
				logs.String("ip", server.Ip),
				logs.Uint16("port", server.Port))
			if err := server.Run(); err != nil {
				a.Close()
				logs.Err("rpc server run fail", logs.String("name", name), logs.Error("err", err))
			}
			ws.Done()
		}(name, s)
	}
	return true
}

func (a *App) NetRpcAllClose() {
	for _, s := range a.rpcServers {
		s.Close()
	}
}

func (a *App) NetRpcClientClose() {
	if a.rpcClientProxy == nil {
		return
	}
	a.rpcClientProxy.Close()
}
