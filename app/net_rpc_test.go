package app

import (
	"context"
	"fmt"
	"testing"

	"github.com/cnlisea/ant/logs"
)

type TestService struct {
	Id int
}

func NewTestService(id int) *TestService {
	return &TestService{Id: id}
}

type TestServiceAddRequest struct {
	A int
	B int
}

type TestServiceAddReply struct {
	C int
}

func (s *TestService) Add(ctx context.Context, args *TestServiceAddRequest, reply *TestServiceAddReply) error {
	logs.Debug("add", logs.Int("id", s.Id), logs.Int("a", args.A), logs.Int("b", args.B))
	reply.C = args.A + args.B
	return nil
}

func TestApp_NetRpcRegister(t *testing.T) {
	var (
		app = New()
		err error
	)
	if err = app.Logger("stdout", logs.LevelDebug, true, 0); err != nil {
		t.Fatal("logger fail", err)
	}

	if err = app.Discovery("public", []*DiscoveryNode{
		{
			Addr: "",
			Port: 8848,
		},
	}); err != nil {
		t.Fatal("discovery fail", err)
	}

	if err = app.NetRpcRegister("test-rpc", "127.0.0.1", 9999, "", nil, NewTestService(0)); err != nil {
		t.Fatal("net rpc register fail", err)
	}
	if err = app.Run(); err != nil {
		t.Fatal("run fail", err)
	}
}

func TestApp_NetRpcRegister2(t *testing.T) {
	var (
		app = New()
		err error
	)
	if err = app.Logger("stdout", logs.LevelDebug, true, 0); err != nil {
		t.Fatal("logger fail", err)
	}

	if err = app.Discovery("public", []*DiscoveryNode{
		{
			Addr: "",
			Port: 8848,
		},
	}); err != nil {
		t.Fatal("discovery fail", err)
	}

	if err = app.NetRpcRegister("test-rpc", "127.0.0.1", 8888, "", nil, NewTestService(1)); err != nil {
		t.Fatal("net rpc register fail", err)
	}
	if err = app.Run(); err != nil {
		t.Fatal("run fail", err)
	}
}

func TestApp_NetRpcRegisterSoftState1(t *testing.T) {
	var (
		app = New()
		err error
	)
	if err = app.Logger("stdout", logs.LevelDebug, true, 0); err != nil {
		t.Fatal("logger fail", err)
	}

	if err = app.Discovery("public", []*DiscoveryNode{
		{
			Addr: "",
			Port: 8848,
		},
	}); err != nil {
		t.Fatal("discovery fail", err)
	}

	discoverySoftState := false
	if err = app.NetRpcRegister("test-rpc", "127.0.0.1", 9999, "", &discoverySoftState, NewTestService(1)); err != nil {
		t.Fatal("net rpc register fail", err)
	}

	if err = app.Run(); err != nil {
		t.Fatal("run fail", err)
	}
}

func TestApp_NetRpcRegisterSoftState2(t *testing.T) {
	var (
		app = New()
		err error
	)
	if err = app.Logger("stdout", logs.LevelDebug, true, 0); err != nil {
		t.Fatal("logger fail", err)
	}

	if err = app.Discovery("public", []*DiscoveryNode{
		{
			Addr: "",
			Port: 8848,
		},
	}); err != nil {
		t.Fatal("discovery fail", err)
	}

	discoverySoftState := true
	if err = app.NetRpcRegister("test-rpc", "127.0.0.2", 8888, "", &discoverySoftState, NewTestService(2)); err != nil {
		t.Fatal("net rpc register fail", err)
	}

	if err = app.Run(); err != nil {
		t.Fatal("run fail", err)
	}
}

func TestApp_NetRpcRegisterSoftState3(t *testing.T) {
	var (
		app = New()
		err error
	)
	if err = app.Logger("stdout", logs.LevelDebug, true, 0); err != nil {
		t.Fatal("logger fail", err)
	}

	if err = app.Discovery("public", []*DiscoveryNode{
		{
			Addr: "",
			Port: 8848,
		},
	}); err != nil {
		t.Fatal("discovery fail", err)
	}

	discoverySoftState := true
	if err = app.NetRpcRegister("test-rpc", "127.0.0.3", 7777, "", &discoverySoftState, NewTestService(7)); err != nil {
		t.Fatal("net rpc register fail", err)
	}

	if err = app.Run(); err != nil {
		t.Fatal("run fail", err)
	}
}

func TestApp_NetRpcClient(t *testing.T) {
	var (
		app = New()
		err error
	)
	if err = app.Logger("", logs.LevelDebug, true, 0); err != nil {
		t.Fatal("logger fail", err)
	}
	if err = app.Discovery("public", []*DiscoveryNode{
		{
			Addr: "",
			Port: 8848,
		},
	}); err != nil {
		t.Fatal("discovery fail", err)
	}
	clientProxy := app.ProxyRpcClient()
	client, err := clientProxy.GetClient("test-rpc", "TestService")
	if err != nil {
		t.Fatal("proxy get client fail", err)
	}
	for i := 0; i < 10; i++ {
		var reply TestServiceAddReply
		if err = client.Invoke(context.Background(), "Add", &TestServiceAddRequest{
			A: (101 - i) * (i + 1),
			B: (402 - 1) * (i + 1),
		}, &reply); err != nil {
			t.Fatal("rpc invoke fail", err)
		}
		t.Log(reply.C)
	}
}

func TestApp_NetRpcClientOneWay(t *testing.T) {
	var (
		app = New()
		err error
	)
	if err = app.Logger("", logs.LevelDebug, true, 0); err != nil {
		t.Fatal("logger fail", err)
	}
	if err = app.Discovery("public", []*DiscoveryNode{
		{
			Addr: "",
			Port: 8848,
		},
	}); err != nil {
		t.Fatal("discovery fail", err)
	}
	clientProxy := app.ProxyRpcClient()
	client, err := clientProxy.GetClient("test-rpc", "TestService")
	if err != nil {
		t.Fatal("proxy get client fail", err)
	}
	if err = client.OneWay(context.Background(), "Add", &TestServiceAddRequest{
		A: 100,
		B: 200,
	}); err != nil {
		t.Fatal("rpc invoke fail", err)
	}
}

func TestApp_NetRpcClientBroadcast(t *testing.T) {
	var (
		app = New()
		err error
	)
	if err = app.Logger("", logs.LevelDebug, true, 0); err != nil {
		t.Fatal("logger fail", err)
	}
	if err = app.Discovery("public", []*DiscoveryNode{
		{
			Addr: "",
			Port: 8848,
		},
	}); err != nil {
		t.Fatal("discovery fail", err)
	}
	clientProxy := app.ProxyRpcClient()
	client, err := clientProxy.GetClient("test-rpc", "TestService")
	if err != nil {
		t.Fatal("proxy get client fail", err)
	}
	var reply TestServiceAddReply
	if err = client.Broadcast(context.Background(), "Add", &TestServiceAddRequest{
		A: 100,
		B: 200,
	}, &reply); err != nil {
		t.Fatal("rpc invoke fail", err)
	}
	fmt.Println(reply.C)
}

func TestApp_NetRpcClientCallback(t *testing.T) {
	var (
		app = New()
		err error
	)
	if err = app.Logger("", logs.LevelDebug, true, 0); err != nil {
		t.Fatal("logger fail", err)
	}
	if err = app.Discovery("public", []*DiscoveryNode{
		{
			Addr: "",
			Port: 8848,
		},
	}); err != nil {
		t.Fatal("discovery fail", err)
	}
	clientProxy := app.ProxyRpcClient()
	client, err := clientProxy.GetClient("test-rpc", "TestService")
	if err != nil {
		t.Fatal("proxy get client fail", err)
	}
	exit := make(chan struct{}, 1)
	client.Callback(context.Background(), "Add", &TestServiceAddRequest{
		A: 100,
		B: 200,
	}, &TestServiceAddReply{}, func(reply interface{}, err error) {
		t.Log(err)
		if rep, ok := reply.(*TestServiceAddReply); ok {
			t.Log(rep.C)
		}
		exit <- struct{}{}
	})
	<-exit
}

func TestApp_NetRpcClientWithSoftState(t *testing.T) {
	var (
		app = New()
		err error
	)
	if err = app.Logger("", logs.LevelDebug, true, 0); err != nil {
		t.Fatal("logger fail", err)
	}
	if err = app.Discovery("public", []*DiscoveryNode{
		{
			Addr: "",
			Port: 8848,
		},
	}); err != nil {
		t.Fatal("discovery fail", err)
	}
	clientProxy := app.ProxyRpcClient()
	client, err := clientProxy.GetClientSoftState("test-rpc", "TestService")
	if err != nil {
		t.Fatal("proxy get client fail", err)
	}
	for i := 0; i < 10; i++ {
		var reply TestServiceAddReply
		if err = client.Invoke(client.WithSoftStateKey(context.Background(), "0"), "Add", &TestServiceAddRequest{
			A: (101 - i) * (i + 1),
			B: (402 - 1) * (i + 1),
		}, &reply); err != nil {
			t.Fatal("rpc invoke fail", err)
		}
		t.Log(reply.C)
	}
}

func TestApp_NetRpcClientWithHash(t *testing.T) {
	var (
		app = New()
		err error
	)
	if err = app.Logger("", logs.LevelDebug, true, 0); err != nil {
		t.Fatal("logger fail", err)
	}
	if err = app.Discovery("public", []*DiscoveryNode{
		{
			Addr: "",
			Port: 8848,
		},
	}); err != nil {
		t.Fatal("discovery fail", err)
	}
	clientProxy := app.ProxyRpcClient()
	client, err := clientProxy.GetClientHash("test-rpc", "TestService")
	if err != nil {
		t.Fatal("proxy get client fail", err)
	}
	for i := 0; i < 10; i++ {
		var reply TestServiceAddReply
		if err = client.Invoke(client.WithHashKey(context.Background(), "127.0.0.3:7777"), "Add", &TestServiceAddRequest{
			A: (101 - i) * (i + 1),
			B: (402 - 1) * (i + 1),
		}, &reply); err != nil {
			t.Fatal("rpc invoke fail", err)
		}
		t.Log(reply.C)
	}
}
