package app

import (
	"context"
	"errors"
	"sync"

	"github.com/cnlisea/ant/app/db/mongo"
	"github.com/cnlisea/ant/app/db/mysql"
	"github.com/cnlisea/ant/app/db/redis"
	"github.com/cnlisea/ant/app/mq"
	netHttp "github.com/cnlisea/ant/app/net/http"
	netRpc "github.com/cnlisea/ant/app/net/rpc"
	"github.com/cnlisea/ant/config"
	"github.com/cnlisea/ant/logs"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
)

type App struct {
	ctx    context.Context
	cancel context.CancelFunc
	closed bool
	mutex  *sync.RWMutex

	cfg                  *config.Config
	discovery            naming_client.INamingClient
	discoveryClusterName string
	discoverySoftState   map[string]bool
	rpcServers           map[string]*netRpc.Server
	rpcServerGroupName   map[string]string
	rpcClientProxy       *netRpc.ClientProxy

	httpServers map[string]*netHttp.Server

	mysql map[string]*mysql.Client
	mongo map[string]*mongo.Client
	redis map[string]*redis.Client

	mqConsumer map[string]mq.Consumer
	mqProducer map[string]mq.Producer
}

func New() *App {
	a := &App{
		mutex: new(sync.RWMutex),
	}
	a.ctx, a.cancel = context.WithCancel(context.Background())
	return a
}

func (a *App) Run() error {
	ws := new(sync.WaitGroup)
	if !a.NetHttpAllStart(ws) ||
		!a.NetRpcAllStart(ws) ||
		!a.MQProducerAllStart(ws) ||
		!a.MQConsumerAllStart(ws) ||
		!a.SignalStart(ws) ||
		!a.DiscoveryStart(ws) {
		logs.Err("service start fail")
		return errors.New("service start fail")
	}
	logs.Info("run successfully...")
	ws.Wait()

	logs.Info("run exit successfully...")
	return nil
}

func (a *App) Context() context.Context {
	return a.ctx
}

func (a *App) Close() {
	a.mutex.Lock()
	if !a.closed {
		a.closed = true
		a.DiscoveryClose()
		a.MQConsumerAllClose()
		a.NetHttpAllClose()
		a.NetRpcAllClose()
		a.NetRpcClientClose()
		a.DBMongoAllClose()
		a.DBMySqlAllClose()
		a.DBRedisAllClose()
		a.MQProducerAllClose()
		a.cancel()
	}
	a.mutex.Unlock()
}
