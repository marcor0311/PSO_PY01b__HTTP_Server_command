package constants

const (
	RouteFibonacci  = "/fibonacci"
	RouteCreateFile = "/createfile"
	RouteDeleteFile = "/deletefile"
	RouteReverse    = "/reverse"
	RouteToUpper    = "/toupper"
	RouteRandom     = "/random"
	RouteTimestamp  = "/timestamp"
	RouteHash       = "/hash"
	RouteSimulate   = "/simulate"
	RouteSleep      = "/sleep"
	RouteLoadTest   = "/loadtest"
	RouteHelp       = "/help"
	RouteStatus     = "/status"
)

const (
	ParallelRouteCount = "/countwords"
	ParallelRouteMatrix = "/matrix"
)

var ParallelRoutes = []string{
    ParallelRouteCount,
    ParallelRouteMatrix,
}

const (
	DispatcherRoutePing = "/ping"
	DispatcherRouteWorkers = "/workers"
)

const (
	StatusOK                  = "200 OK"
	StatusBadRequest          = "400 Bad Request"
	StatusNotFound            = "404 Not Found"
	StatusInternalServerError = "500 Internal Server Error"
	StatusServiceUnavailable  = "503 Unavailable"
	StatusBadGateway		  = "502 Bad Gateway"
)

const (
	DISPATCHER = "Dispatcher"
	WORKER     = "Worker"
)