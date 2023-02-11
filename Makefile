gen:
	protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. proto/*.proto

clean:
	rm **/*.pb.go 

server:
	go run cmd/server/main.go

client:
	go run cmd/client/main.go
