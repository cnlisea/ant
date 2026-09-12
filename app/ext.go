package app

func (a *App) ExtGet(key string) (any, bool) {
	return a.ext.Load(key)
}

func (a *App) ExtSet(key string, val any) {
	a.ext.Store(key, val)
}
