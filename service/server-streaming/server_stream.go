package serverstreaming

import (
	"fmt"

	"github.com/ric-v/grpc-example/pb"
)

type ServerStream struct {
}

func (s *ServerStream) ServerStream(req *pb.ServerStreamRequest, stream pb.ServerStream_ServerStreamServer) error {
	for i := 0; i < 10; i++ {
		if err := stream.Send(&pb.ServerStreamRequest{Name: req.Name}); err != nil {
			break
		}
	}
	fmt.Println("ServerStream done")
	return nil
}
