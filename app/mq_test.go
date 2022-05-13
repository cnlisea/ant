package app

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/cnlisea/ant/app/mq/message"
	"github.com/cnlisea/ant/logs"
)

func TestApp_MQConsumerRegister(t *testing.T) {
	var (
		app = New()
		err error
	)
	if err = app.Logger("", logs.LevelDebug, true, 0); err != nil {
		t.Fatal(err)
	}

	if err = app.MQConsumerRegister("", "", "", []string{""}, "", "GID_TEST_CONSUMER", false, 32, []*MQConsumerSubscribe{
		{
			Topic: "test",
			Tag:   "",
			Handler: func(ctx context.Context, consumer ...*message.Consumer) bool {
				for i := range consumer {
					t.Log("index:", i+1, "data:", string(consumer[i].Data))
				}
				return true
			},
		},
	}); err != nil {
		t.Fatal("mq consumer register fail", err)
	}

	if err = app.Run(); err != nil {
		t.Fatal("run fail", err)
	}
}

func TestApp_MQProducerRegister(t *testing.T) {
	var (
		app = New()
		err error
	)
	if err = app.Logger("", logs.LevelDebug, true, 0); err != nil {
		t.Fatal(err)
	}

	if err = app.MQProducerRegister("", "", "", []string{""}, "", "GID_TEST_PRODUCER"); err != nil {
		t.Fatal("mq producer register fail", err)
	}

	go func() {
		time.Sleep(1 * time.Second)
		var (
			proxy = app.ProxyMQ()
			ctx   = context.Background()
		)
		for i := 0; i < 10; i++ {
			if err = proxy.SendMsg(ctx, "", "test", []byte(fmt.Sprintf("index:%d", i+1))); err != nil {
				t.Error("send msg", err)
			}
		}
		t.Log("success")
		time.Sleep(3 * time.Second)
		app.Close()
	}()

	if err = app.Run(); err != nil {
		t.Fatal("run fail", err)
	}

}
