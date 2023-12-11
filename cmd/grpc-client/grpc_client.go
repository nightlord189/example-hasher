package main

import (
	"context"
	"fmt"
	grpc2 "github.com/nightlord189/example-hasher/internal/delivery/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial("127.0.0.1:5300", opts...)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	client := grpc2.NewHasherClient(conn)
	request := &grpc2.HashRequest{
		Items: []*grpc2.HashRequestItem{
			{
				Id:   "1",
				Data: "test_string",
				Type: 1,
			},
			{
				Id:   "2",
				Data: "Lorem Ipsum",
				Type: 1,
			},
		},
	}
	response, err := client.GetHashes(context.Background(), request)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	fmt.Println(response.Message, response.Items)
}
