package main

import (
	"log"
	"net"

	"github.com/cory-evans/monorepo-example/pkg/person"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	person.RegisterService(s)

	log.Printf("server is listening on port :8080")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
