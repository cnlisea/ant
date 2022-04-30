package app

type _DiscoverySoftState = string

const (
	_DiscoverySoftStateRpcNamePrefix  = _DiscoverySoftState("rpc_")
	_DiscoverySoftStateHttpNamePrefix = _DiscoverySoftState("http_")
)

func (a *App) _DiscoverySoftStateAdd(prefix _DiscoverySoftState, name string, state bool) {
	if a.discoverySoftState == nil {
		a.discoverySoftState = make(map[string]bool, 1)
	}
	a.discoverySoftState[prefix+name] = state
}

func (a *App) _DiscoverySoftState(prefix _DiscoverySoftState, name string) (state bool, exist bool) {
	if a.discoverySoftState == nil {
		return
	}

	state, exist = a.discoverySoftState[prefix+name]
	return
}
