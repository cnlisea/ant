package config

import (
	"container/list"

	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
)

type Config struct {
	client config_client.IConfigClient
	units  *list.List
}

func New() *Config {
	return new(Config)
}
