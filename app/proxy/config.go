package proxy

type Config interface {
	GetCfg(key ...string) interface{}
}
