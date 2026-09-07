package proxy

import (
	"github.com/cnlisea/ant/logs"
	"net/http"
)

type Config interface {
	GetCfg(key ...string) interface{}
	SetCfg(key string, obj interface{})
	DiscoveryRegMethod() func(
		namespaceId string,
		nodes []*DiscoveryNode) error
	NetHttpRegMethod() func(
		name string,
		ip string,
		port uint16,
		discoverySoftState *bool,
		handler http.Handler) error
	DBMongoRegMethod() func(
		name string,
		user string,
		password string,
		addr []string,
		dbName string,
		replicaSet string,
		connTimeout int,
		active int, idle int,
		idleTimeout int) error
	DBMySQLRegMethod() func(
		name string,
		user string,
		password string,
		addr string,
		port uint16,
		dbName string,
		charset string,
		connTimeout string,
		parseTime bool,
		loc string,
		active int,
		idle int,
		idleTimeout int) error
	DBRedisRegMethod() func(
		name string,
		password string,
		addr string,
		port uint16,
		db int,
		active int,
		idle int,
		idleTimeout int) error
	MQProducerRegMethod() func(
		name string,
		accessKey string,
		secretKey string,
		nameServer []string,
		namespace string,
		groupId string) error
	MQConsumerRegMethod() func(
		name string,
		accessKey string,
		secretKey string,
		nameServer []string,
		namespace string,
		groupId string,
		broadCastModel bool,
		batchSize int,
		subscribes []*MQConsumerSubscribe) error
	LogRegMethod() func(
		path string,
		level logs.Level,
		json bool,
		callerSkip int) error
}
