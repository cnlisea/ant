package config

import (
	"errors"
	"fmt"

	"github.com/cnlisea/ant/config/layout"
	"github.com/cnlisea/ant/logs"
	"github.com/cnlisea/ant/typex"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

type Register struct {
	Key       string
	Name      string
	Path      string
	Obj       interface{}
	Layout    Layout
	LoadLocal bool
	Center    *RegisterCenter
}

type RegisterCenter struct {
	GroupId    string
	DataId     string
	UpdateHook func(obj interface{})
}

func (c *Config) Register(reg *Register) error {
	if reg == nil {
		return errors.New("param is nil")
	}

	if reg.Center != nil && c.client == nil {
		return errors.New("unset remote")
	}

	var lt layout.Layout
	switch reg.Layout {
	case LayoutYaml:
		lt = new(layout.Yaml)
	case LayoutJson:
		lt = new(layout.Json)
	case LayoutToml:
		lt = new(layout.Toml)
	case LayoutIni:
		lt = new(layout.Ini)
	default:
		return errors.New("invalid layout")
	}

	unit := &Unit{
		key:    reg.Key,
		name:   reg.Name,
		path:   reg.Path,
		obj:    reg.Obj,
		Layout: lt,
	}
	if c.UnitsExist(unit, c.UnitsEqual) {
		return fmt.Errorf("config already existed key: %s, name:%s, path:%s", reg.Key, reg.Name, reg.Path)
	}

	var err error
	if reg.LoadLocal {
		err = unit.LoadLocal()
		if err != nil {
			return fmt.Errorf("config load local fail key: %s, name:%s, path:%s", reg.Key, reg.Name, reg.Path)
		}
	}

	if reg.Center != nil {
		unit.updateHook = reg.Center.UpdateHook
		if err = c.client.ListenConfig(vo.ConfigParam{
			DataId: reg.Center.DataId,
			Group:  reg.Center.GroupId,
			OnChange: func(namespace, group, dataId, data string) {
				if err = unit.Update(typex.StringToBytes(data)); err != nil {
					logs.Warn("unit update fail", logs.String("err", err.Error()))
				}
			},
		}); err != nil {
			return fmt.Errorf("config listen fail, err:%v", err)
		}

		var data string
		data, err = c.client.GetConfig(vo.ConfigParam{
			DataId: reg.Center.DataId,
			Group:  reg.Center.GroupId,
		})
		if err != nil {
			return fmt.Errorf("config get fail, err:%v", err)
		}

		if err = unit.Update(typex.StringToBytes(data)); err != nil {
			return fmt.Errorf("unit update fail, err:%v", err)
		}
	}
	c.UnitsAdd(unit)

	return nil
}
