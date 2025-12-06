package main

import (
    "log"
    "net"
    "google.golang.org/grpc"
    
    pb "github.com/stepan41k/GinTest/service-1/pkg/api/v1"
    "github.com/stepan41k/GinTest/service-1/internal/app" 
)

func main() {
    lis, _ := net.Listen("tcp", ":50051")
    s := grpc.NewServer()

    srv := app.NewGRPCServer() 
    
    pb.RegisterGreeterServer(s, srv)

    if err := s.Serve(lis); err != nil {
        log.Fatal(err)
    }
}