package main

import (
	"fmt"
	"log"
	"time"

	"github.com/simonvetter/modbus"
)

func main() {
	client, err := modbus.NewClient(&modbus.ClientConfiguration{
		URL:     "tcp://192.168.1.50:502",
		Timeout: 2 * time.Second,
		Speed:   19200,
	})

	if err != nil {
		log.Fatal(err)
	}

	if err := client.Open(); err != nil {
		log.Fatalf("Не удалось подключиться: %v", err)
	}
	defer client.Close()

	registers, err := client.ReadRegisters(100, 2, modbus.HOLDING_REGISTER)
	if err != nil {
		log.Fatalf("Ошибка чтения: %v", err)
	}

	fmt.Printf("Значения регистров: %v\n", registers)
}
