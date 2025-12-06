package main

import (
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	service_b "github.com/stepan41k/GinTest/service-2/internal/app"
	pb "github.com/stepan41k/GinTest/service-1/pkg/api/v1"
)

func main() {
	const grpcAddress = "localhost:50051"
	
	conn, err := grpc.NewClient(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Не удалось подключиться к gRPC серверу: %v", err)
	}
	
	defer conn.Close()

	greeterClient := pb.NewGreeterClient(conn)

	app := service_b.NewApp(greeterClient)

	http.HandleFunc("/hello", app.HandleHello)

	log.Println("Сервис B (HTTP) запущен на :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Ошибка сервера: %v", err)
	}
}