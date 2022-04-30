package app

import (
	"net/http"
	"sync"

	netHttp "github.com/cnlisea/ant/app/net/http"
	"github.com/cnlisea/ant/logs"
)

func (a *App) NetHttpRegister(name string, ip string, port uint16, discoverySoftState *bool, handler http.Handler) error {
	if a.httpServers == nil {
		a.httpServers = make(map[string]*netHttp.Server, 1)
	}
	s := a.httpServers[name]
	if s == nil {
		s = netHttp.NewServer(ip, port)
		a.httpServers[name] = s
	}

	if discoverySoftState != nil {
		a._DiscoverySoftStateAdd(_DiscoverySoftStateHttpNamePrefix, name, *discoverySoftState)
	}
	return s.SetHandler(handler)
}

func (a *App) NetHttpAllStart(ws *sync.WaitGroup) bool {
	var (
		name string
		s    *netHttp.Server
	)
	for name, s = range a.httpServers {
		ws.Add(1)
		go func(name string, server *netHttp.Server) {
			logs.Info("http server run",
				logs.String("name", name),
				logs.String("ip", s.Ip),
				logs.Uint16("port", s.Port))
			if err := server.Run(); err != nil {
				a.Close()
				logs.Err("http server run fail", logs.String("name", name), logs.Error("err", err))
			}
			ws.Done()
		}(name, s)
	}
	return true
}

func (a *App) NetHttpAllClose() {
	for _, s := range a.httpServers {
		s.Close()
	}
}
