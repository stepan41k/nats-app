package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {
	// 1. Подключение к NATS
	nc, err := nats.Connect("nats://127.0.0.1:4222")
	if err != nil {
		log.Fatalf("Ошибка подключения: %s", err.Error())
	}
	defer nc.Drain()

	// 2. Инициализация JetStream контекста
	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatalf("Ошибка инициализации контекста: %s", err.Error())
	}

	ctx := context.Background()

	// 3. Создание (или получение) Stream
	// Stream сохраняет сообщения. Без него сообщения уйдут в никуда.
	_, err = js.CreateOrUpdateStream(ctx, jetstream.StreamConfig{
		Name:     "ORDERS",
		Subjects: []string{"orders.>"}, // Слушаем все темы начинающиеся с orders.
	})
	if err != nil {
		log.Fatalf("Ощибка создания stream: %s", err.Error())
	}

	// 4. Публикация сообщений
	for i := 1; i <= 5; i++ {
		msg := fmt.Sprintf("Order Payload #%d", i)

		// Публикуем асинхронно для скорости, но можно и синхронно (Publish)
		ack, err := js.Publish(ctx, "orders.new", []byte(msg))
		if err != nil {
			log.Printf("Ошибка публикации: %v", err)
			continue
		}

		log.Printf("Опубликовано: %s (seq: %d)", msg, ack.Sequence)
		time.Sleep(500 * time.Millisecond)
	}
}
