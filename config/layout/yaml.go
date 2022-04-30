package layout

import (
	"gopkg.in/yaml.v2"
)

type Yaml struct{}

func (y *Yaml) ExtName() string {
	return ".yaml"
}

func (y *Yaml) Parse(data []byte, obj interface{}) error {
	return yaml.Unmarshal(data, obj)
}
