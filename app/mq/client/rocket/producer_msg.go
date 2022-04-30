package rocket

import (
	"context"
	"errors"

	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/cnlisea/ant/app/mq/message"
)

func (p *Producer) SendMsg(ctx context.Context, topic string, msg *message.Producer) error {
	if msg == nil {
		return nil
	}

	sendMsg := primitive.NewMessage(topic, msg.Data)
	if msg.ShardingKey != "" {
		sendMsg.WithShardingKey(msg.ShardingKey)
	}
	result, err := p.client.SendSync(ctx, sendMsg)
	if err != nil {
		return err
	}

	if result.Status != primitive.SendOK {
		return errors.New(result.String())
	}

	return nil
}
