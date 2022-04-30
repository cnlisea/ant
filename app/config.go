package app

import (
	"github.com/cnlisea/ant/config"
)

type ConfigCenterNote = config.RemoteNode

func (a *App) ConfigCenter(namespaceId string, nodes []*ConfigCenterNote) error {
	if a.cfg == nil {
		a.cfg = config.New()
	}
	return a.cfg.SetRemote(namespaceId, nodes)
}

type RegisterCenter = config.RegisterCenter

func (a *App) ConfigRegister(key string, name string, path string, loadLocal bool, obj interface{}, rc *RegisterCenter) error {
	if a.cfg == nil {
		a.cfg = config.New()
	}
	return a.cfg.Register(&config.Register{
		Key:       key,
		Name:      name,
		Path:      path,
		LoadLocal: loadLocal,
		Obj:       obj,
		Center:    rc,
	})
}

func (a *App) Config(key string) interface{} {
	if a.cfg == nil {
		return nil
	}
	return a.cfg.GetObj(key)
}
