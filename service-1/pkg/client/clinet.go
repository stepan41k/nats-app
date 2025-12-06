package client

import (
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
    pb "github.com/stepan41k/GinTest/service-1/pkg/api/v1"
)

// NewGreeterClient — удобная обертка для создания клиента
func NewGreeterClient(addr string) (pb.GreeterClient, func(), error) {
    conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        return nil, nil, err
    }
    
    // Возвращаем клиента и функцию закрытия соединения
    return pb.NewGreeterClient(conn), func() { conn.Close() }, nil
}