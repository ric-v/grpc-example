package bidistreaming

import (
	"fmt"
	"io"
	"log"

	"github.com/ric-v/grpc-example/pb"
)

type BidiStream struct {
}

// service BidiStreamService { rpc BidiStream(stream BidiStreamRequest) returns (stream BidiStreamResponse); }
func (s *BidiStream) BiDirectionalStream(stream pb.BiDirectionalStreamService_BiDirectionalStreamServer) error {

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		log.Println("Received request: ", req)

		// send a response
		err = stream.Send(&pb.BiDirectionalStreamResponse{
			Message: fmt.Sprintf("Hello %s", req.Name),
		})
		if err != nil {
			return err
		}

	}
}
