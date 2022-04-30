package app

import (
	"context"
	"sync"

	"github.com/cnlisea/ant/app/mq"
	"github.com/cnlisea/ant/app/mq/message"
	"github.com/cnlisea/ant/app/mq/option"
	"github.com/cnlisea/ant/logs"
)

type MQConsumerSubscribe struct {
	Topic   string
	Tag     string
	Handler func(context.Context, ...*message.Consumer) bool
}

func (a *App) MQConsumerRegister(name string, nameServer []string, namespace string, groupId string, broadCastModel bool, batchSize int, subscribes []*MQConsumerSubscribe) error {
	subscribesLen := len(subscribes)
	if subscribesLen == 0 {
		return nil
	}

	ctx := a.Context()
	client, err := mq.NewConsumerRocket(ctx, nameServer, namespace, groupId, option.ConsumerWithBroadCastModel(broadCastModel), option.ConsumerWithBatchSize(batchSize))
	if err != nil {
		return err
	}

	for i := 0; i < subscribesLen; i++ {
		if err = client.Subscribe(ctx, subscribes[i].Topic, subscribes[i].Tag, subscribes[i].Handler); err != nil {
			return err
		}
	}

	if a.mqConsumer == nil {
		a.mqConsumer = make(map[string]mq.Consumer, 1)
	}
	a.mqConsumer[name] = client
	return nil
}

func (a *App) MQConsumerInstance(name ...string) mq.Consumer {
	if a.mqConsumer == nil {
		return nil
	}

	if name != nil && len(name) > 0 {
		return a.mqConsumer[name[0]]
	}

	return a.mqConsumer[""]
}

func (a *App) MQConsumerAllStart(ws *sync.WaitGroup) bool {
	var (
		ctx  = a.Context()
		name string
		c    mq.Consumer
		err  error
	)
	for name, c = range a.mqConsumer {
		if err = c.Start(ctx); err != nil {
			a.Close()
			logs.Err("mq consumer start fail", logs.String("name", name), logs.Error("err", err))
			return false
		}
	}
	return true
}

func (a *App) MQConsumerAllClose() {
	ctx := a.Context()
	for _, c := range a.mqConsumer {
		c.Close(ctx)
	}
}

func (a *App) MQProducerRegister(name string, accessKey string, secretKey string, nameServer []string, namespace string, groupId string) error {
	client, err := mq.NewProducerRocket(a.ctx, nameServer, namespace, groupId, option.ProducerWithAuth(&option.Auth{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}))
	if err != nil {
		return err
	}

	if a.mqProducer == nil {
		a.mqProducer = make(map[string]mq.Producer, 1)
	}
	a.mqProducer[name] = client
	return nil
}

func (a *App) MQProducerInstance(name ...string) mq.Producer {
	if a.mqProducer == nil {
		return nil
	}

	if name != nil && len(name) > 0 {
		return a.mqProducer[name[0]]
	}

	return a.mqProducer[""]
}

func (a *App) MQProducerAllStart(ws *sync.WaitGroup) bool {
	var (
		ctx  = a.Context()
		name string
		c    mq.Producer
		err  error
	)
	for name, c = range a.mqProducer {
		if err = c.Start(ctx); err != nil {
			a.Close()
			logs.Err("mq producer start fail", logs.String("name", name), logs.Error("err", err))
			return false
		}
	}
	return true
}

func (a *App) MQProducerAllClose() {
	ctx := a.Context()
	for _, c := range a.mqProducer {
		c.Close(ctx)
	}
}
