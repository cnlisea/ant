package config

import (
	"io/ioutil"

	"github.com/cnlisea/ant/config/layout"
)

type Unit struct {
	key        string
	name       string
	path       string
	obj        interface{}
	Layout     layout.Layout
	updateHook func(obj interface{})
	sign       []byte
}

func (u *Unit) LoadLocal() error {
	var name = u.name
	if name == "" {
		name = "dev"
	}

	data, err := ioutil.ReadFile(u.path + name + u.Layout.ExtName())
	if err != nil {
		return err
	}

	return u.UpdateObj(data)
}
