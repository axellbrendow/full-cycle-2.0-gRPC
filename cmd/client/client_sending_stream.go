package main

import (
	"context"
	"fmt"
	"log"
	"time"

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

	AddUsers(client)
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "1",
			Name:  "axell1",
			Email: "axell1@gmail.com",
		},
		&pb.User{
			Id:    "2",
			Name:  "axell2",
			Email: "axell2@gmail.com",
		},
		&pb.User{
			Id:    "3",
			Name:  "axell3",
			Email: "axell3@gmail.com",
		},
		&pb.User{
			Id:    "4",
			Name:  "axell4",
			Email: "axell4@gmail.com",
		},
		&pb.User{
			Id:    "5",
			Name:  "axell5",
			Email: "axell5@gmail.com",
		},
	}

	stream, err := client.AddUsers(context.Background())

	if err != nil {
		log.Fatalf("Could not send data via stream: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)
}
