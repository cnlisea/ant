package config

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
)

type RemoteNode struct {
	Addr string
	Port uint16
}

func (c *Config) SetRemote(namespaceId string, nodes []*RemoteNode) error {
	var (
		nodesLen      = len(nodes)
		serverConfigs = make([]constant.ServerConfig, nodesLen)
		i             int
	)
	for i = 0; i < nodesLen; i++ {
		serverConfigs[i].IpAddr = nodes[i].Addr
		serverConfigs[i].Port = uint64(nodes[i].Port)
	}
	client, err := clients.CreateConfigClient(map[string]interface{}{
		constant.KEY_CLIENT_CONFIG: constant.ClientConfig{
			TimeoutMs:   10 * 1000,
			NamespaceId: namespaceId,
			CacheDir:    "./cache",
			LogDir:      "./log",
		},
		constant.KEY_SERVER_CONFIGS: serverConfigs,
	})
	if err != nil {
		return err
	}

	c.client = client
	return nil
}
