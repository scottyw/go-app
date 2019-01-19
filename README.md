# go-app
An example of a Go application that does nothing in particular

## Build

### Prerequisites

Install Go version 1.11 or later and protobufs:

    brew install go
    brew install protobuf

Make sure your GOPATH is set explicitly:

    export GOPATH=~/go

Install Go tools:

go get -u github.com/jteeuwen/go-bindata/...
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go

### Building

rm -rf generated
mkdir -p generated
protoc \
  -I/usr/local/include \
  -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:generated \
  --swagger_out=logtostderr=true:generated \
  --grpc-gateway_out=logtostderr=true:generated \
  --proto_path proto app.proto
go generate ./...
