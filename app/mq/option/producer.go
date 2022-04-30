package option

type Producer struct {
	Auth *Auth
}

type ProducerFunc func(*Producer)

func ProducerWithAuth(auth *Auth) ProducerFunc {
	return func(option *Producer) {
		option.Auth = auth
	}
}
