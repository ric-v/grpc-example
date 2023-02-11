package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/ric-v/grpc-example/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	c, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer c.Close()

	cli := pb.NewSimpleServiceClient(c)

	req := &pb.SimpleRequest{
		Name:      "Ricardo",
		Tags:      []string{"golang", "grpc"},
		Msg:       "Hello World",
		IdOrSsid:  &pb.SimpleRequest_Id{Id: 123},
		Height:    160.5,
		Weight:    77.2,
		Age:       30,
		Timestamp: time.Now().Unix(),
		IsOk:      true,
		Data:      []byte("Hello World"),
	}

	rsp, err := cli.Hello(context.Background(), req)
	if err != nil {
		panic(err)
	}
	fmt.Println("Hello : ", rsp)

	rsp, err = cli.Info(context.Background(), req)
	if err != nil {
		panic(err)
	}
	fmt.Println("Info : ", rsp)

	echoResp, err := cli.Echo(context.Background(), req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Echo : %+v\n", echoResp)

	// Server streaming
	cli2 := pb.NewServerStreamClient(c)

	stream, err := cli2.ServerStream(context.Background(), &pb.ServerStreamRequest{Name: "Ricardo"})
	if err != nil {
		panic(err)
	}

	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}
		fmt.Printf("Server streaming : %+v\n", resp)
	}

	// Client streaming
	cli3 := pb.NewUploadServiceClient(c)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	stream2, err := cli3.Upload(ctx)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 3; i++ {
		if err := stream2.Send(&pb.UploadRequest{
			Filename: "file" + strconv.Itoa(i) + ".txt",
			Data:     []byte("Hello World " + strconv.Itoa(i)),
		}); err != nil {
			panic(err)
		}
	}

	resp, err := stream2.CloseAndRecv()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Client streaming : %+v\n", resp)

	// Bidi streaming
	cli4 := pb.NewBiDirectionalStreamServiceClient(c)

	stream3, err := cli4.BiDirectionalStream(context.Background())
	if err != nil {
		panic(err)
	}

	go func() {

		for i := 0; i < 3; i++ {
			if err := stream3.Send(&pb.BiDirectionalStreamRequest{
				Name: "Ricardo",
			}); err != nil {
				panic(err)
			}
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			resp, err := stream3.Recv()
			if err != nil {
				break
			}
			fmt.Printf("Bidi streaming : %+v\n", resp)
		}
	}()

	time.Sleep(time.Second * 10)
}
