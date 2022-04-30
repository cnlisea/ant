package typex

import (
	"io"

	"github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func JsonMarshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func JsonUnmarshalReader(reader io.Reader, v interface{}) error {
	return json.NewDecoder(reader).Decode(v)
}

func JsonUnmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
