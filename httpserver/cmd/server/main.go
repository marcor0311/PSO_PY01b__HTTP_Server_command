package main

import (
	"log"

	"httpserver/internal/tcp"
)

func main() {
	server, err := tcp.CreateTcpClient(":8080")
	if err != nil {
		log.Fatalf("Failed to start TCP server: %v", err)
	}

	for msg := range server.ReceiveChan {
		log.Printf("[Received] Method: %s, Path: %s\n", msg.Method, msg.Path)
	}
}
