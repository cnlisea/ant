package app

type ProxyExt struct {
	a *App
}

func (pe *ProxyExt) Get(key string) (any, bool) {
	return pe.a.ExtGet(key)
}

func (pe *ProxyExt) Set(key string, val any) {
	pe.a.ExtSet(key, val)
}
