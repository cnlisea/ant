package message

type Producer struct {
	ShardingKey string
	Data        []byte
}
