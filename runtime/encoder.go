package runtime

import (
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type MessageEncoder interface {
	Encode(x interface{}) ([]byte, error)
	Decode(data []byte, x interface{}) error
}

type JSONEncoder struct{}

func (JSONEncoder) Encode(x interface{}) ([]byte, error) {
	return json.Marshal(x)
}

func (JSONEncoder) Decode(data []byte, x interface{}) error {
	return json.Unmarshal(data, x)
}

type ProtobufEncoder struct{}

func (ProtobufEncoder) Encode(x interface{}) ([]byte, error) {
	msg, ok := x.(protoreflect.ProtoMessage)
	if !ok {
		return nil, fmt.Errorf("x must be protoreflect.ProtoMessage")
	}
	return proto.Marshal(msg)
}

func (ProtobufEncoder) Decode(data []byte, x interface{}) error {
	msg, ok := x.(protoreflect.ProtoMessage)
	if !ok {
		return fmt.Errorf("x must be protoreflect.ProtoMessage")
	}
	return proto.Unmarshal(data, msg)
}
