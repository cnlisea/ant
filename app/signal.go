package app

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/cnlisea/ant/logs"
)

func (a *App) Signal() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGPIPE,
		syscall.SIGHUP,
	)

	var sig os.Signal
Loop:
	for {
		select {
		case sig = <-sc:
			switch sig {
			case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				logs.Info("Main got signal", logs.String("signal", sig.String()))
				a.Close()
				break Loop
			case syscall.SIGPIPE:
				logs.Info("Main Ignore broken pipe signal")
			case syscall.SIGHUP:
				logs.Info("Main Got update config signal")
				/*
					if err := a.Reload(); err != nil {
						logs.Error("server config reload fail")
					}
				*/
			}
		case <-a.ctx.Done():
			break Loop
		}
	}
}

func (a *App) SignalStart(ws *sync.WaitGroup) bool {
	ws.Add(1)
	go func() {
		a.Signal()
		ws.Done()
	}()
	return true
}
