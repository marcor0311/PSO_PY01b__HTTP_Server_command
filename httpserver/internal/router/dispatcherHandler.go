package router

import (
	"net"

	"httpserver/internal/constants"
	"httpserver/internal/utils"
)

// Ping
func handlePing(conn net.Conn, path string) {
	utils.WriteHTTPResponse(conn, constants.StatusOK, "pong")
}

// Workers
func handleWorkers(conn net.Conn, path string) {
	utils.WriteHTTPResponse(conn, constants.StatusOK, "pong")
}