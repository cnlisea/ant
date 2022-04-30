package app

import (
	"testing"

	"github.com/cnlisea/ant/logs"
)

func TestApp_Logger(t *testing.T) {
	var (
		app = New()
		err error
	)
	if err = app.Logger("stdout", logs.LevelDebug, true, 0); err != nil {
		t.Fatal(err)
	}

	logs.Debug("app logger debug")
	logs.Info("app logger info")
	logs.Warn("app logger warn")
	logs.Err("app logger err")
	logs.Fatal("app logger fatal")
}
