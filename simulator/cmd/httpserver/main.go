package main

import (
	"fmt"
	"simulator/internal/httpserver/handlers"
)

func main() {
	fmt.Println("Hora actual del sistema:", handlers.Timestamp())
}