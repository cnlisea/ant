package app

import (
	"context"
	"database/sql"

	"github.com/cnlisea/ant/app/mq/message"
	"github.com/cnlisea/ant/app/proxy"
	"github.com/gomodule/redigo/redis"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProxyConfig struct {
	a *App
}

func (pc *ProxyConfig) GetCfg(key ...string) interface{} {
	var k string
	if key != nil && len(key) > 0 {
		k = key[0]
	}
	return pc.a.Config(k)
}

func (a *App) ProxyConfig() proxy.Config {
	return &ProxyConfig{
		a: a,
	}
}

type ProxyRpcClientPool struct {
	a         *App
	groupName string
}

func (prc *ProxyRpcClientPool) GetClient(name string, serviceName string) (proxy.RpcClient, error) {
	return prc.a.NetRpcClient(prc.groupName, name, serviceName)
}

func (prc *ProxyRpcClientPool) GetClientHash(name string, serviceName string) (proxy.RpcClient, error) {
	return prc.a.NetRpcClient(prc.groupName, name, serviceName, prc.a.NetRpcClientWithHashSelector())
}

func (prc *ProxyRpcClientPool) GetClientSoftState(name string, serviceName string) (proxy.RpcClient, error) {
	return prc.a.NetRpcClient(prc.groupName, name, serviceName, prc.a.NetRpcClientWithSoftState())
}

func (a *App) ProxyRpcClient(groupName ...string) proxy.RpcClientPool {
	var name string
	if groupName != nil && len(groupName) > 0 {
		name = groupName[0]
	}
	return &ProxyRpcClientPool{
		a:         a,
		groupName: name,
	}
}

type ProxyDBMySQL struct {
	a *App
}

func (pdm *ProxyDBMySQL) GetDB(name ...string) *sql.DB {
	instance := pdm.a.DBMySqlInstance(name...)
	if instance == nil {
		return nil
	}
	return instance.GetConn(pdm.a.Context())
}

func (a *App) ProxyDBMySQL() proxy.DBMySQL {
	return &ProxyDBMySQL{
		a: a,
	}
}

type ProxyDBRedis struct {
	a *App
}

func (pdr *ProxyDBRedis) GetDB(name ...string) redis.Conn {
	instance := pdr.a.DBRedisInstance(name...)
	if instance == nil {
		return nil
	}
	return instance.GetConn(pdr.a.Context())
}

func (a *App) ProxyDBRedis() proxy.DBRedis {
	return &ProxyDBRedis{
		a: a,
	}
}

type ProxyDBMongo struct {
	a *App
}

func (pdm *ProxyDBMongo) GetDB(name ...string) *mongo.Database {
	instance := pdm.a.DBMongoInstance(name...)
	if instance == nil {
		return nil
	}
	return instance.GetConn(pdm.a.Context())
}

func (a *App) ProxyDBMongo() proxy.DBMongo {
	return &ProxyDBMongo{
		a: a,
	}
}

type ProxyMQ struct {
	a *App
}

func (pm *ProxyMQ) SendMsg(ctx context.Context, name string, topic string, msg []byte) error {
	instance := pm.a.MQProducerInstance(name)
	if instance == nil {
		return nil
	}
	return instance.SendMsg(ctx, topic, &message.Producer{
		Data: msg,
	})
}

func (pm *ProxyMQ) SendMsgSharding(ctx context.Context, name string, key string, topic string, msg []byte) error {
	instance := pm.a.MQProducerInstance(name)
	if instance == nil {
		return nil
	}
	return instance.SendMsg(ctx, topic, &message.Producer{
		ShardingKey: key,
		Data:        msg,
	})
}

func (a *App) ProxyMQ() proxy.MQ {
	return &ProxyMQ{
		a: a,
	}
}

type ProxyDiscovery struct {
	a *App
}

func (pd *ProxyDiscovery) Instance() naming_client.INamingClient {
	return pd.a.discovery
}

func (pd *ProxyDiscovery) ClusterName() string {
	return pd.a.DiscoveryClusterName()
}

func (a *App) ProxyDiscovery() proxy.Discovery {
	return &ProxyDiscovery{
		a: a,
	}
}
