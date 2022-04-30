package rocket

import (
	"context"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/cnlisea/ant/app/mq/option"
)

type Consumer struct {
	client rocketmq.PushConsumer
}

func NewConsumer(ctx context.Context, nameServer []string, namespace string, groupId string, options ...option.ConsumerFunc) (*Consumer, error) {
	op := new(option.Consumer)
	for _, f := range options {
		f(op)
	}

	var model = consumer.Clustering
	if op.BroadCastModel {
		model = consumer.BroadCasting
	}

	if op.BatchSize <= 0 {
		op.BatchSize = 1
	}

	var (
		accessKey string
		secretKey string
	)
	if op.Auth != nil {
		accessKey = op.Auth.AccessKey
		secretKey = op.Auth.SecretKey
	}

	c, err := rocketmq.NewPushConsumer(
		//consumer.WithInstance(strconv.FormatUint(uint64(s.cfg.Listen.Port), 10)),
		consumer.WithConsumerModel(model),
		consumer.WithNameServer(nameServer),
		consumer.WithNamespace(namespace),
		consumer.WithGroupName(groupId),
		consumer.WithConsumeMessageBatchMaxSize(op.BatchSize),
		consumer.WithCredentials(primitive.Credentials{
			AccessKey: accessKey,
			SecretKey: secretKey,
		}))
	if err != nil {
		return nil, err
	}

	return &Consumer{
		client: c,
	}, nil
}

func (c *Consumer) Start(ctx context.Context) error {
	return c.client.Start()
}

func (c *Consumer) Close(ctx context.Context) error {
	return c.client.Shutdown()
}
