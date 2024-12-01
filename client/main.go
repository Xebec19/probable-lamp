package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/Xebec19/probable-lamp/greeting"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var name = flag.String("name", "Default Name", "Name to greet")

func main() {

	flag.Parse()

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server at localhost:50051: %v", err)
	}

	defer conn.Close()

	c := pb.NewGreetingServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayGreeting(ctx, &pb.GreetingRequest{Name: *name})
	if err != nil {
		log.Fatalf("error calling function SayGreeting: %v", err)
	}

	log.Printf("Response from gRPC server's SayGreeting function: %s", r.GetMessage())
}
