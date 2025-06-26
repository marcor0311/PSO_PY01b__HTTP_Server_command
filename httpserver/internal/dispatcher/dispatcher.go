package dispatcher

import (
	"fmt"
	"log"
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
	const maxAttempts = 2

	var lastErr error

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		worker := worker.ChooseWorker()
		if worker == nil {
			utils.WriteHTTPResponse(connection, constants.StatusServiceUnavailable,
				"There are no available workers")
			return
		}

		log.Printf("[Dispatcher] Attempt %d: forwarding to worker %s", attempt, worker.Address)

		targetURL := worker.Address + path
		req, err := http.NewRequest(method, targetURL, nil)
		if err != nil {
			lastErr = err
			utils.WriteHTTPResponse(connection, constants.StatusBadRequest,
				fmt.Sprintf("[Dispatcher] Request build error: %v", err))
			return
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Printf("[Dispatcher] Error contacting %s", worker.Address)
			worker.MarkInactive()
			lastErr = err
			continue
		}
		defer resp.Body.Close()

		if err := utils.CopyHTTPResponse(connection, resp); err != nil {
			log.Printf("[Dispatcher] Copy response error: %v", err)
		}

		worker.Completed++
		worker.Load++
		return
	}

	utils.WriteHTTPResponse(connection, constants.StatusBadGateway,
		fmt.Sprintf("[Dispatcher] All forwarding attempts failed: %v", lastErr))
}
