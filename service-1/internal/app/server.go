package app

import (
    "context"
    pb "github.com/stepan41k/GinTest/service-1/pkg/api/v1"
)

type GRPCServer struct {
    pb.UnimplementedGreeterServer
}

func NewGRPCServer() *GRPCServer {
    return &GRPCServer{}
}

// Реализация метода
func (s *GRPCServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
    return &pb.HelloResponse{Message: "Hello " + req.Name}, nil
}