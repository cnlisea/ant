package option

type Consumer struct {
	Id             string
	BroadCastModel bool
	BatchSize      int
	Auth           *Auth
}

type ConsumerFunc func(*Consumer)

func ConsumerWithId(id string) ConsumerFunc {
	return func(option *Consumer) {
		option.Id = id
	}
}

func ConsumerWithBroadCastModel(broadCastModel bool) ConsumerFunc {
	return func(option *Consumer) {
		option.BroadCastModel = broadCastModel
	}
}

func ConsumerWithBatchSize(batchSize int) ConsumerFunc {
	return func(option *Consumer) {
		option.BatchSize = batchSize
	}
}

func ConsumerWithAuth(auth *Auth) ConsumerFunc {
	return func(option *Consumer) {
		option.Auth = auth
	}
}
