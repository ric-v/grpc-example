package main

import (
	"net"

	"github.com/ric-v/grpc-example/pb"
	clientstreaming "github.com/ric-v/grpc-example/service/client-streaming"
	serverstreaming "github.com/ric-v/grpc-example/service/server-streaming"
	bidistreaming "github.com/ric-v/grpc-example/service/bidi-streaming"
	"github.com/ric-v/grpc-example/service/unary"
	"google.golang.org/grpc"
)

func main() {

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	// add unary service (unary.UnaryService) and server streaming service (serverstreaming.ServerStream)
	// to the grpc server

	s := grpc.NewServer()

	pb.RegisterSimpleServiceServer(s, &unary.UnaryService{})
	pb.RegisterServerStreamServer(s, &serverstreaming.ServerStream{})
	pb.RegisterUploadServiceServer(s, &clientstreaming.ClientStream{})
	pb.RegisterBiDirectionalStreamServiceServer(s, &bidistreaming.BidiStream{})

	s.Serve(l)
}
