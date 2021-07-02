package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/marcelovbm/go-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	// AddUser(client)
	// AddUserVerbose(client)
	// AddUsers(client)
	AddUserStreamBoth(client)
}

func AddUser(client pb.UserServiceClient) {

	request := &pb.User{
		Id:    "0",
		Name:  "Marcelo Magrinelli",
		Email: "marcelovbm@gmail.com",
	}

	response, err := client.AddUser(context.Background(), request)

	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	fmt.Println(response)
}

func AddUserVerbose(client pb.UserServiceClient) {
	request := &pb.User{
		Id:    "0",
		Name:  "Marcelo Magrinelli",
		Email: "marcelovbm@gmail.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), request)

	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not receive the message: %v", err)
		}

		fmt.Println("Status: ", stream.Status, " - ", stream.GetUser())
	}
}

func AddUsers(client pb.UserServiceClient) {
	requests := []*pb.User{
		&pb.User{
			Id:    "0",
			Name:  "Marcelo Magrinelli",
			Email: "marcelovbm@gmail.com",
		}, &pb.User{
			Id:    "1",
			Name:  "Marcelo Magrinelli",
			Email: "marcelovbm@gmail.com",
		}, &pb.User{
			Id:    "2",
			Name:  "Marcelo Magrinelli",
			Email: "marcelovbm@gmail.com",
		}, &pb.User{
			Id:    "3",
			Name:  "Marcelo Magrinelli",
			Email: "marcelovbm@gmail.com",
		}, &pb.User{
			Id:    "4",
			Name:  "Marcelo Magrinelli",
			Email: "marcelovbm@gmail.com",
		}, &pb.User{
			Id:    "5",
			Name:  "Marcelo Magrinelli",
			Email: "marcelovbm@gmail.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, request := range requests {
		stream.Send(request)
		time.Sleep(time.Second * 3)
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(response)
}

func AddUserStreamBoth(client pb.UserServiceClient) {

	stream, err := client.AddUserStreamBoth(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	requests := []*pb.User{
		&pb.User{
			Id:    "0",
			Name:  "Marcelo Magrinelli",
			Email: "marcelovbm@gmail.com",
		}, &pb.User{
			Id:    "1",
			Name:  "Marcelo Magrinelli",
			Email: "marcelovbm@gmail.com",
		}, &pb.User{
			Id:    "2",
			Name:  "Marcelo Magrinelli",
			Email: "marcelovbm@gmail.com",
		}, &pb.User{
			Id:    "3",
			Name:  "Marcelo Magrinelli",
			Email: "marcelovbm@gmail.com",
		}, &pb.User{
			Id:    "4",
			Name:  "Marcelo Magrinelli",
			Email: "marcelovbm@gmail.com",
		}, &pb.User{
			Id:    "5",
			Name:  "Marcelo Magrinelli",
			Email: "marcelovbm@gmail.com",
		},
	}

	wait := make(chan int)

	go func() {
		for _, request := range requests {
			fmt.Println("(Client) Sending user: ", request.GetName())
			stream.Send(request)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			response, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("(Client) Erro receving data: ", err)
				break
			}
			fmt.Printf("(Client) Receiving user %v com status: %v\n", response.GetUser().GetName(), response.GetStatus())
		}
		close(wait)
	}()

	<-wait
}
