package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/backendservice/Study-gRPC/grpcService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct {
}

func (s *server) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserReplay, error) {
	return &pb.GetUserReplay{UserId: "user1"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterDataEngineServer(s, &server{})

	fmt.Println("Server started ...")

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
