package router

import (
	"net"
	"strings"

	"httpserver/internal/constants"
)

type RouteHandler func(conn net.Conn, path string)

var routes = map[string]RouteHandler{}

func init() {
	routes[constants.RouteFibonacci] = handleFibonacci
	routes[constants.RouteCreateFile] = handleCreateFile
	routes[constants.RouteDeleteFile] = handleDeleteFile
	routes[constants.RouteReverse] = handleReverse
	routes[constants.RouteToUpper] = handleToUpper
	routes[constants.RouteRandom] = handleRandom
	routes[constants.RouteTimestamp] = handleTimestamp
	routes[constants.RouteHash] = handleHash
	routes[constants.RouteSimulate] = handleSimulate
	routes[constants.RouteSleep] = handleSleep
	routes[constants.RouteLoadTest] = handleLoadTest
	routes[constants.RouteHelp] = handleHelp
	routes[constants.RouteStatus] = handleStatus
}

func HandleRoute(path string, conn net.Conn) {
	cleanPath := strings.SplitN(path, "?", 2)[0]
	if handler, exists := routes[cleanPath]; exists {
		handler(conn, path)
	} else {
		handleNotFound(conn, path)
	}
}