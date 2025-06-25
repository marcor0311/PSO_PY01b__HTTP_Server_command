package dispatcher

import (
	"fmt"
	"net"
	"net/http"

	"httpserver/internal/constants"
	"httpserver/internal/utils"
	"httpserver/internal/worker"
)

/* Forwards the incoming request to an available worker and relays the worker’s response back to the client.
 *
 * @param {string} method - HTTP method.
 * @param {string} path - Requested path and query string.
 * @param {net.Conn} connection - TCP connection to the original client.
 */
func Forward(method, path string, connection net.Conn) {
	worker := worker.ChooseWorker()
	if worker == nil {
		utils.WriteHTTPResponse(connection, constants.StatusServiceUnavailable,
			"There are not available workers")
		return
	}

	targetURL := worker.Address + path

	req, err := http.NewRequest(method, targetURL, nil)
	if err != nil {
		utils.WriteHTTPResponse(connection, constants.StatusBadRequest,
			fmt.Sprintf("[Dispatcher] Request build error: %v", err))
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		utils.WriteHTTPResponse(connection, constants.StatusBadGateway,
			fmt.Sprintf("[Dispatcher] Worker error: %v", err))
		return
	}
	defer resp.Body.Close()

	if err := utils.CopyHTTPResponse(connection, resp); err != nil {
		fmt.Printf("[Dispatcher] Copy response error: %v\n", err)
		return
	}

	worker.Completed++
	worker.Load++
}
