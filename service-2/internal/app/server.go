package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	pb "github.com/stepan41k/GinTest/service-1/pkg/api/v1"
)

type App struct {
	grpcClient pb.GreeterClient
}

func NewApp(client pb.GreeterClient) *App {
	return &App{
		grpcClient: client,
	}
}

func (a *App) HandleHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	resp, err := a.grpcClient.SayHello(ctx, &pb.HelloRequest{Name: name})
	
	if err != nil {
		http.Error(w, "Ошибка при вызове gRPC сервиса: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Сервис B получил ответ от Сервиса A: %s", resp.GetMessage())
}