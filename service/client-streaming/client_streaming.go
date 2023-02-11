package clientstreaming

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"

	"github.com/ric-v/grpc-example/pb"
)

type ClientStream struct {
}

// service UploadService { rpc Upload(UploadRequest) returns (UploadResponse); }

func (s *ClientStream) Upload(stream pb.UploadService_UploadServer) error {

	var res []*pb.FileInfo
	for {

		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}

		// do something with the first request
		err = os.WriteFile("tmp/"+req.Filename, req.Data, 0644)
		if err != nil {
			return err
		}

		// get the hash of the file
		hash, err := hashFile(req.Filename)
		if err != nil {
			return err
		}
		fmt.Println(hash)

		// create a response
		res = append(res, &pb.FileInfo{
			Filename: req.Filename,
		})
	}
	fmt.Print(res)

	// send the response
	err := stream.SendAndClose(&pb.UploadResponse{
		Files: res,
	})
	if err != nil {
		return err
	}

	return nil
}

func hashFile(filename string) (string, error) {

	// get the hash of the file
	f, _ := os.ReadFile(filename)
	return string(sha256.New().Sum(f)), nil
}
