package app

import (
	"strconv"
	"sync"

	netRpc "github.com/cnlisea/ant/app/net/rpc"
	"github.com/cnlisea/ant/logs"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

type DiscoveryNode struct {
	Addr string
	Port uint16
}

func (a *App) Discovery(namespaceId string, nodes []*DiscoveryNode) error {
	var (
		nodeLen      = len(nodes)
		serverConfig = make([]constant.ServerConfig, nodeLen)
		i            int
	)
	for i = 0; i < nodeLen; i++ {
		serverConfig[i].IpAddr = nodes[i].Addr
		serverConfig[i].Port = uint64(nodes[i].Port)
	}
	client, err := clients.CreateNamingClient(map[string]interface{}{
		constant.KEY_CLIENT_CONFIG: constant.ClientConfig{
			TimeoutMs:            10 * 1000,
			BeatInterval:         5 * 1000,
			NamespaceId:          namespaceId,
			CacheDir:             "./cache",
			LogDir:               "./log",
			UpdateThreadNum:      20,
			NotLoadCacheAtStart:  true,
			UpdateCacheWhenEmpty: false,
		},
		constant.KEY_SERVER_CONFIGS: serverConfig,
	})
	if err != nil {
		return err
	}

	a.discovery = client
	return nil
}

func (a *App) DiscoveryClusterName() string {
	return "public"
}

func (a *App) DiscoveryRegister() error {
	if a.discovery == nil {
		return ErrDiscoveryUnavailable
	}

	var (
		name         string
		s            *netRpc.Server
		ip           string
		port         uint16
		groupName    string
		metaData     map[string]string
		state, exist bool
		err          error
	)
	for name, s = range a.rpcServers {
		ip, port = s.Addr()
		metaData = map[string]string{"network": "tcp"}
		state, exist = a._DiscoverySoftState(_DiscoverySoftStateRpcNamePrefix, name)
		if exist {
			metaData["offline"] = strconv.FormatBool(!state)
		}
		groupName = ""
		if a.rpcServerGroupName != nil {
			groupName = a.rpcServerGroupName[name]
		}
		if groupName == "" {
			groupName = "DEFAULT_GROUP"
		}
		if _, err = a.discovery.RegisterInstance(vo.RegisterInstanceParam{
			Ip:          ip,
			Port:        uint64(port),
			Metadata:    metaData,
			GroupName:   groupName,
			ServiceName: name,
			Weight:      10,
			ClusterName: a.DiscoveryClusterName(),
			Enable:      true,
			Healthy:     true,
			Ephemeral:   true,
		}); err != nil {
			logs.Err("discovery register fail",
				logs.String("name", name),
				logs.String("ip", ip),
				logs.Uint16("port", port),
				logs.Error("err", err))
			return err
		}
		logs.Info("discovery register success",
			logs.String("name", name),
			logs.String("ip", ip),
			logs.Uint16("port", port))
	}

	return err
}

func (a *App) DiscoveryCancel() error {
	if a.discovery == nil {
		return ErrDiscoveryUnavailable
	}

	var (
		name      string
		s         *netRpc.Server
		ip        string
		port      uint16
		groupName string
	)
	for name, s = range a.rpcServers {
		ip, port = s.Addr()
		groupName = ""
		if a.rpcServerGroupName != nil {
			groupName = a.rpcServerGroupName[name]
		}
		if groupName == "" {
			groupName = "DEFAULT_GROUP"
		}
		a.discovery.DeregisterInstance(vo.DeregisterInstanceParam{
			Ip:          ip,
			Port:        uint64(port),
			ServiceName: name,
			Cluster:     a.DiscoveryClusterName(),
			GroupName:   groupName,
			Ephemeral:   true,
		})
	}
	return nil
}

func (a *App) DiscoveryStart(ws *sync.WaitGroup) bool {
	if err := a.DiscoveryRegister(); err != nil && err != ErrDiscoveryUnavailable {
		a.Close()
		return false
	}
	return true
}

func (a *App) DiscoveryClose() {
	a.DiscoveryCancel()
}
