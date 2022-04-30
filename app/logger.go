package app

import (
	"github.com/cnlisea/ant/logs"
)

func (a *App) Logger(path string, level logs.Level, json bool, callerSkip int) error {
	if path == "" {
		path = "stdout"
	}
	return logs.Init(path, level, json, callerSkip)
}
