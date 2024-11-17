package main

import (
	"context"
	"log"
	"net"

	pb "github.com/Xebec19/probable-lamp/greeting"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreetingServiceServer
}

func (s *server) SayGreeting(ctx context.Context, in *pb.GreetingRequest) (*pb.GreetingResponse, error) {
	return &pb.GreetingResponse{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreetingServiceServer(s, &server{})
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
