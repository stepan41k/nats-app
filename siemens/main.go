package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"time"

	"github.com/robinson/gos7"
)

func main() {
	// Rack 0, Slot 1 - стандартно для S7-1200/1500 (для S7-300 обычно Slot 2)
	handler := gos7.NewTCPClientHandler("192.168.1.10", 0, 1)
	handler.Timeout = 5 * time.Second
	handler.IdleTimeout = 5 * time.Second

	if err := handler.Connect(); err != nil {
		log.Fatalf("Ошибка подключения: %v", err)
	}
	defer handler.Close()

	client := gos7.NewClient(handler)

	buf := make([]byte, 10)
	err := client.AGReadDB(1, 0, len(buf), buf)
	if err != nil {
		log.Fatalf("Ошибка чтения DB: %v", err)
	}

	var helper gos7.Helper
	realVal := helper.GetRealAt(buf, 0) // Значение REAL из DB1.DBD0
	intVal := int16(binary.BigEndian.Uint16(buf[4:6]))    // Значение INT из DB1.DBW4

	fmt.Printf("Real: %f, Int: %d\n", realVal, intVal)
}
