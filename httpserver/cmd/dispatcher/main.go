package main

import (
	"fmt"
	"httpserver/internal/constants"
	"httpserver/internal/tcp"
	"httpserver/internal/utils"
	"httpserver/internal/worker"
	"log"
)

func main() {
	worker.RegisterWorkersFromEnv()

	port := utils.GetEnv("PORT", "8080")
	serverPort := fmt.Sprintf(":%s", port)

	server, err := tcp.CreateTcpClient(serverPort, constants.DISPATCHER)
	if err != nil {
		log.Fatalf("[Dispatcher] Failed to start TCP server: %v", err)
	}

	for msg := range server.ReceiveChan {
		log.Printf("[Dispatcher] Method: %s, Path: %s\n", msg.Method, msg.Path)
	}
}
