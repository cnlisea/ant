package proxy

import (
	"github.com/cnlisea/ant/logs"
)

type ConfigDiscoveryNode struct {
	Addr string
	Port uint16
}

type Config interface {
	GetCfg(key ...string) interface{}
	DiscoveryRegMethod() func(namespaceId string, nodes []*DiscoveryNode) error
	DBMongoRegMethod() func(name string,
		user string, password string,
		addr []string, dbName string,
		replicaSet string, connTimeout int,
		active int, idle int, idleTimeout int) error
	DBMySQLRegMethod() func(name string,
		user string, password string,
		addr string, port uint16, dbName string,
		charset string, connTimeout string, parseTime bool, loc string,
		active int, idle int, idleTimeout int) error
	DBRedisRegMethod() func(name string,
		password string, addr string, port uint16, db int,
		active int, idle int, idleTimeout int) error
	LogRegMethod() func(path string, level logs.Level, json bool, callerSkip int) error
}
