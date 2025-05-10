package main

import (
	"fmt"
	"httpserver/internal/handlers"
)

func main() {
	fmt.Println("Hora actual del sistema:", handlers.Timestamp())
}