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
	RoutePing		= "/ping"
)

const (
	StatusOK                  = "200 OK"
	StatusBadRequest          = "400 Bad Request"
	StatusNotFound            = "404 Not Found"
	StatusInternalServerError = "500 Internal Server Error"
)

const (
	DISPATCHER      = "Dispatcher"
	WORKER          = "Worker"
)