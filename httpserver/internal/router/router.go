package router

import (
	"bufio"
	"net"
	"strings"

	"httpserver/internal/constants"
)

type RouteHandler func(net.Conn, string, *bufio.Reader)

var routes = map[string]RouteHandler{}
var dispatcherRoutes = map[string]RouteHandler{}

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
	routes[constants.ParallelRouteCount] = handleWordCountChunk
	routes[constants.ParallelRouteMontecarlo] = handleMontecarlo
	dispatcherRoutes[constants.DispatcherRouteWorkers] = handleWorkers
	dispatcherRoutes[constants.DispatcherRoutePing] = handlePing
	dispatcherRoutes[constants.ParallelRouteCount] = handlParallelWordCount
	dispatcherRoutes[constants.ParallelRouteMontecarlo] = handleParallelMontecarlo
}

func HandleRoute(path string, conn net.Conn, br *bufio.Reader) {
	cleanPath := strings.SplitN(path, "?", 2)[0]
	if handler, exists := routes[cleanPath]; exists {
		handler(conn, path, br)
	} else {
		handleNotFound(conn, path, br)
	}
}

func HandleDispatcherRouter(path string, conn net.Conn, br *bufio.Reader) bool {
	cleanPath := strings.SplitN(path, "?", 2)[0]
	if handler, exists := dispatcherRoutes[cleanPath]; exists {
		handler(conn, path, br)
		return true
	}
	return false
}
