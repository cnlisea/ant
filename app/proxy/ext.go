package proxy

type Ext interface {
	Get(key string) (any, bool)
	Set(key string, ext any)
}
