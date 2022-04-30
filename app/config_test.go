package app

import (
	"testing"

	"github.com/cnlisea/ant/logs"
)

func TestApp_ConfigCenter(t *testing.T) {
	var (
		app = New()
		err error
	)
	if err = app.Logger("", logs.LevelDebug, true, 0); err != nil {
		t.Fatal("logger fail", err)
	}
	if err = app.ConfigCenter("public", []*ConfigCenterNote{
		{
			Addr: "",
			Port: 8848,
		},
	}); err != nil {
		t.Fatal("config center fail", err)
	}
	t.Log("config center successfully!!!")
}

func TestApp_ConfigRegister(t *testing.T) {
	var (
		app = New()
		err error
	)
	if err = app.Logger("", logs.LevelDebug, true, 0); err != nil {
		t.Fatal("logger fail", err)
	}

	var obj struct {
		Addr string
		Port uint32
	}
	if err = app.ConfigRegister("", "dev", "./", true, &obj, nil); err != nil {
		t.Fatal("config register fail", err)
	}
	t.Log(obj)
}

func TestApp_ConfigCenterRegister(t *testing.T) {
	var (
		app = New()
		err error
	)
	if err = app.Logger("", logs.LevelDebug, true, 0); err != nil {
		t.Fatal("logger fail", err)
	}
	t.Log("logger successfully!!!")

	if err = app.ConfigCenter("public", []*ConfigCenterNote{
		{
			Addr: "",
			Port: 8848,
		},
	}); err != nil {
		t.Fatal("config center fail", err)
	}
	t.Log("config center successfully!!!")

	var obj struct {
		Addr string
		Port uint32
	}
	if err = app.ConfigRegister("", "dev", "./", false, &obj, &RegisterCenter{
		GroupId: "test",
		DataId:  "test",
		UpdateHook: func(obj interface{}) {
			t.Log("config update", obj)
		},
	}); err != nil {
		t.Fatal("config register fail", err)
	}
	t.Log(obj)
	select {}
}

func TestApp_ConfigLoadLocalCenterRegister(t *testing.T) {
	var (
		app = New()
		err error
	)
	if err = app.Logger("", logs.LevelDebug, true, 0); err != nil {
		t.Fatal("logger fail", err)
	}
	t.Log("logger successfully!!!")

	if err = app.ConfigCenter("public", []*ConfigCenterNote{
		{
			Addr: "",
			Port: 8848,
		},
	}); err != nil {
		t.Fatal("config center fail", err)
	}
	t.Log("config center successfully!!!")

	var obj struct {
		Addr string
		Port uint32
	}
	if err = app.ConfigRegister("", "dev", "./", true, &obj, &RegisterCenter{
		GroupId: "test",
		DataId:  "test",
		UpdateHook: func(obj interface{}) {
			t.Log("config update", obj)
		},
	}); err != nil {
		t.Fatal("config register fail", err)
	}
	t.Log(obj)
	select {}
}
