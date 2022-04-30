package layout

import (
	"github.com/pelletier/go-toml/v2"
)

type Toml struct{}

func (t *Toml) ExtName() string {
	return ".toml"
}

func (t *Toml) Parse(data []byte, obj interface{}) error {
	return toml.Unmarshal(data, obj)
}
