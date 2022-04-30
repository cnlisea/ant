package rocket

import (
	"context"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/cnlisea/ant/app/mq/option"
)

type Producer struct {
	client rocketmq.Producer
}

func NewProducer(ctx context.Context, nameServer []string, namespace string, groupId string, options ...option.ProducerFunc) (*Producer, error) {
	op := new(option.Producer)
	for _, f := range options {
		f(op)
	}
	var accessKey, secretKey string
	if op.Auth != nil {
		accessKey = op.Auth.AccessKey
		secretKey = op.Auth.SecretKey
	}
	client, err := rocketmq.NewProducer(
		producer.WithNameServer(nameServer),
		producer.WithNamespace(namespace),
		producer.WithGroupName(groupId),
		producer.WithQueueSelector(NewProducerHashSelector()),
		producer.WithCredentials(primitive.Credentials{
			AccessKey: accessKey,
			SecretKey: secretKey,
		}))
	if err != nil {
		return nil, err
	}

	return &Producer{
		client: client,
	}, nil
}

func (p *Producer) Start(ctx context.Context) error {
	return p.client.Start()
}

func (p *Producer) Close(ctx context.Context) error {
	return p.client.Shutdown()
}
