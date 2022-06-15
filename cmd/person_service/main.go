package main

import (
	"context"
	"log"
	"net"

	"github.com/cory-evans/monorepo-example/pkg/proto/person"
	"google.golang.org/grpc"
)

type server struct {
	person.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, req *person.HelloRequest) (*person.HelloResponse, error) {
	msg := "Hello, " + req.Name

	return &person.HelloResponse{Message: msg}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	person.RegisterHelloServiceServer(s, &server{})

	log.Printf("server is listening on port :8080")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
