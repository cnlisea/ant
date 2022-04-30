package jwt

type Jwt struct {
	Key []byte
}

func New(key []byte) *Jwt {
	return &Jwt{
		Key: key,
	}
}
