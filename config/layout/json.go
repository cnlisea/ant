package layout

import (
	"github.com/cnlisea/ant/typex"
)

type Json struct{}

func (j *Json) ExtName() string {
	return ".json"
}

func (j *Json) Parse(data []byte, obj interface{}) error {
	return typex.JsonUnmarshal(data, obj)
}
