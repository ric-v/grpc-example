package unary

import (
	"context"

	"github.com/ric-v/grpc-example/pb"
)

type UnaryService struct {
}

func New() *UnaryService {
	return &UnaryService{}
}

func (s *UnaryService) Hello(ctx context.Context, req *pb.SimpleRequest) (rsp *pb.SimpleResponse, err error) {
	return &pb.SimpleResponse{
		Ack:  true,
		Msg:  "Hello " + req.Name,
		Code: 200,
	}, nil
}

func (s *UnaryService) Echo(ctx context.Context, req *pb.SimpleRequest) (rsp *pb.SimpleRequest, err error) {
	return req, nil
}

func (s *UnaryService) Info(ctx context.Context, req *pb.SimpleRequest) (rsp *pb.SimpleResponse, err error) {
	return &pb.SimpleResponse{
		Ack:  true,
		Msg:  "Info " + req.Name,
		Code: 200,
	}, nil
}
