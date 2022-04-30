package proxy

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
)

type Discovery interface {
	Instance() naming_client.INamingClient
	ClusterName() string
}
