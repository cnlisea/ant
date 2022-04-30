package rocket

import (
	"strconv"

	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

type ProducerHashSelector struct {
	RoundRobin producer.QueueSelector
}

func NewProducerHashSelector() *ProducerHashSelector {
	return &ProducerHashSelector{
		RoundRobin: producer.NewRoundRobinQueueSelector(),
	}
}

func (m *ProducerHashSelector) Select(message *primitive.Message, queues []*primitive.MessageQueue) *primitive.MessageQueue {
	v := message.GetShardingKey()
	if v == "" {
		return m.RoundRobin.Select(message, queues)
	}
	message.RemoveProperty(primitive.PropertyShardingKey)

	key, err := strconv.Atoi(v)
	if err != nil || key < 0 {
		return m.RoundRobin.Select(message, queues)
	}

	return queues[key%len(queues)]
}
