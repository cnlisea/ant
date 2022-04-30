package rocket

import (
	"context"

	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/cnlisea/ant/app/mq/message"
)

func (c *Consumer) Subscribe(ctx context.Context, topic string, tag string, f func(context.Context, ...*message.Consumer) bool) error {
	if tag == "" {
		tag = "*"
	}
	return c.client.Subscribe(topic, consumer.MessageSelector{
		Type:       consumer.TAG,
		Expression: tag,
	}, func(ctx context.Context, msg ...*primitive.MessageExt) (result consumer.ConsumeResult, err error) {
		msgLen := len(msg)
		if msgLen == 0 {
			return consumer.ConsumeSuccess, nil
		}

		var (
			fMsg = make([]*message.Consumer, msgLen)
			i    int
		)
		for i = 0; i < msgLen; i++ {
			fMsg[i] = &message.Consumer{
				Data: msg[i].Body,
			}
		}
		result = consumer.ConsumeSuccess
		if !f(ctx, fMsg...) {
			result = consumer.ConsumeRetryLater
		}
		return result, nil
	})
}

func (c *Consumer) Unsubscribe(ctx context.Context, topic string) error {
	return c.client.Unsubscribe(topic)
}
