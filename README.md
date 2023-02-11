# GRPC Example

## Install protoc
```bash
# https://github.com/protocolbuffers/protobuf/releases & add to path OR
brew install protobuf
```

## Golang plugins
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

export PATH="$PATH:$(go env GOPATH)/bin"
```

## Init go module
```bash
go mod init github.com/ric-v/grpc-example
```

## Generate gRPC code
```bash
protoc --go_out=plugins=grpc:. \
    --proto_path=proto \
     proto/*.proto
```

