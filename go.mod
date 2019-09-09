module github.com/tresko/shippy-service-consignment

go 1.13

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1

require (
	github.com/golang/protobuf v1.3.2
	github.com/micro/go-micro v1.9.1
	github.com/micro/protoc-gen-micro v0.8.0 // indirect
	github.com/nats-io/nats-server/v2 v2.0.4 // indirect
	golang.org/x/net v0.0.0-20190909003024-a7b16738d86b
)
