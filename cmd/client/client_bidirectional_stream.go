package main

import (
	"context"
	"fmt"
	"io"
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

	AddUserBidirectionalStream(client)
}

func AddUserBidirectionalStream(client pb.UserServiceClient) {
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

	stream, err := client.AddUserBidirectionalStream(context.Background())

	if err != nil {
		log.Fatalf("Could not get stream: %v", err)
	}

	wait := make(chan int)

	go func() {
		for _, req := range reqs {
			fmt.Println("Sending user: ", req.Name)
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Could not receive the stream message: %v", err)
				break
			}

			fmt.Printf("Receiving user %v with status: %v\n", res.GetUser().GetName(), res.GetStatus())
		}
		close(wait)
	}()

	<-wait
}
