package person

import (
	"context"

	pb "github.com/cory-evans/monorepo-example/pkg/proto/person"
	"google.golang.org/grpc"
)

type Service struct {
	pb.UnimplementedHelloServiceServer
}

func (s *Service) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	msg := "Hello, " + req.Name

	return &pb.HelloResponse{Message: msg}, nil
}

func RegisterService(s *grpc.Server) {
	pb.RegisterHelloServiceServer(s, &Service{})
}
