package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/ua"
)

func main() {
	ctx := context.Background()
	endpoint := "opc.tcp://192.168.1.100:4840"

	c, err := opcua.NewClient(endpoint, opcua.SecurityMode(ua.MessageSecurityModeNone))
	if err != nil {
		log.Fatalf("Ошибка создания клиента: %v", err)
	}

	if err := c.Connect(ctx); err != nil {
		log.Fatalf("Ошибка подключения: %v", err)
	}

	defer c.Close(ctx)

	nodeID, err := ua.ParseNodeID("ns=2;s=MyVariable")
	if err != nil {
		log.Fatal(err)
	}

	req := &ua.ReadRequest{
		NodesToRead: []*ua.ReadValueID{
			{NodeID: nodeID},
		},
	}

	resp, err := c.Read(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Results[0].Status == ua.StatusOK {
		fmt.Printf("Значение: %v\n", resp.Results[0].Value.Value())
	} else {
		fmt.Printf("Ошибка чтения: %v\n", resp.Results[0].Status)
	}
}
