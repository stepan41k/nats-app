package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {
	// 1. Подключение к NATS
	nc, err := nats.Connect("nats://127.0.0.1:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Drain()

	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// 2. Получение доступа к Stream
	stream, err := js.Stream(ctx, "ORDERS")
	if err != nil {
		log.Fatal("Stream не найден. Сначала запустите producer:", err)
	}

	// 3. Создание (или получение) Durable Consumer
	// Durable означает, что сервер запомнит, какие сообщения мы уже обработали,
	// даже если мы отключимся.
	consumer, err := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Durable:   "OrderProcessor", // Уникальное имя консьюмера
		AckPolicy: jetstream.AckExplicitPolicy, // Мы будем явно подтверждать обработку
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Консьюмер запущен. Ожидание сообщений...")

	// 4. Обработка сообщений (Consume)
	// Метод Consume автоматически "вытягивает" сообщения батчами
	consContext, err := consumer.Consume(func(msg jetstream.Msg) {
		fmt.Printf("Получено сообщение: %s\n", string(msg.Data()))

		// Имитация работы
		// time.Sleep(100 * time.Millisecond)

		// ВАЖНО: Подтверждаем сообщение, иначе оно придет снова
		msg.Ack()
	})
	if err != nil {
		log.Fatal(err)
	}
	defer consContext.Stop()

	// Ожидание сигнала завершения (Ctrl+C)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	fmt.Println("\nЗавершение работы...")
}