package main

import (
	"fmt"
	"log"

	"httpserver/internal/constants"
	"httpserver/internal/tcp"
)

func main() {
	const port = "8080"
	listenAddr := fmt.Sprintf(":%s", port)

	server, err := tcp.CreateTcpClient(listenAddr, constants.WORKER)
	if err != nil {
		log.Fatalf("Failed to start TCP server: %v", err)
	}

	for msg := range server.ReceiveChan {
		log.Printf("[Received] Method: %s, Path: %s\n", msg.Method, msg.Path)
	}
}
