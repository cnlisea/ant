package layout

import (
	"gopkg.in/ini.v1"
)

type Ini struct{}

func (i *Ini) ExtName() string {
	return ".ini"
}

func (i *Ini) Parse(data []byte, obj interface{}) error {
	return ini.MapTo(obj, data)
}
