package router

import (
	"net"
	"strings"

	"httpserver/internal/constants"
)

var dispatcherRoutes = map[string]RouteHandler{}

func init() {
	dispatcherRoutes[constants.DispatcherRouteWorkers] = handleFibonacci
	dispatcherRoutes[constants.DispatcherRoutePing] = handlePing
}

func HandleDispatcherRouter(path string, conn net.Conn) {
	cleanPath := strings.SplitN(path, "?", 2)[0]
	if handler, exists := routes[cleanPath]; exists {
		handler(conn, path)
	}
}
