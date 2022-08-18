# protoc-gen-broker is a protoc plugin to generate a broker interface for protobuf messages.

# Description
The goal is generating contracts by protobuf models. That can makes easy data transfer with broker. Plugin supports JSON and protobuf encodings.

The plugin generates a broker interface with methods:
- method for sending to broker;
- method for handling received message.

# Build
Execute below command for building protoc plugin.
```sh
make build
```
protoc plugin has been made in bin folder.

# Example

There is content [example.proto](example/example.proto).
```
syntax = "proto3";

package example;

import "broker/broker.proto";

option go_package = "example/proto";

message Foo {
    option (broker.generate) = true;
    option (broker.encoder) = ENCODER_JSON;

    string example = 1;
}

message Bar {
    option (broker.generate) = true;
    option (broker.encoder) = ENCODER_PROTOBUF;

    string example = 1;
}

message FooBar {
    string example = 1;
}
```

Option __broker.generate__ uses for generating a broker interface. Option __broker.encoder__ allows select an encode method (default __JSON__)

Next, generate __pb__ files from example.proto using below command:
```sh
PATH=`pwd`/bin:$PATH protoc -I . --proto_path=example \
    --go_out=. --go_opt=paths=source_relative \
    --broker_out=. --broker_opt=paths=source_relative \
    example/example.proto
```

Messages __Foo__ and __Bar__ have broker interfaces. You can see it in file __example/example_broker.pb.go__:
```golang
...
// Foo broker interface
type FooBrokerProducerInterface interface {
	Write(context.Context, *Foo) error
}
type FooHandler func(context.Context, *Foo) error
type FooBrokerConsumerInterface interface {
	Handle(context.Context, *Foo) (io.Closer, error)
}

// Bar broker interface
type BarBrokerProducerInterface interface {
	Write(context.Context, *Bar) error
}
type BarHandler func(context.Context, *Bar) error
type BarBrokerConsumerInterface interface {
	Handle(context.Context, *Bar) (io.Closer, error)
}
...
```

# Warning
You can generate broker interface in only folder with generated protobuf models now. Maybe I will fix it in near future.

# License
[MIT](LICENSE)
