package typex

import (
	"github.com/gogo/protobuf/proto"
)

func ProtoMarshal(pb proto.Message) ([]byte, error) {
	return proto.Marshal(pb)
}

func ProtoUnmarshal(buf []byte, pb proto.Message) error {
	return proto.Unmarshal(buf, pb)
}
