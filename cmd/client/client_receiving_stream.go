package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/axell-brendow/full-cycle-2.0-gRPC/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	AddUserVerbose(client)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "axell",
		Email: "axell@gmail.com",
	}

	stream, err := client.AddUserVerbose(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Could not receive the stream message: %v", err)
			break
		}

		fmt.Println("Status: ", res.Status, " - ", res.GetUser())
	}
}
