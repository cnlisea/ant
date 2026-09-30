package typex

import (
	"github.com/gogo/protobuf/proto"
)

type ProtoMessage = proto.Message

func ProtoMarshal(pb ProtoMessage) ([]byte, error) {
	return proto.Marshal(pb)
}

func ProtoUnmarshal(buf []byte, pb ProtoMessage) error {
	return proto.Unmarshal(buf, pb)
}
