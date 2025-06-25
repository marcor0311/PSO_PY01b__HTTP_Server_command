package dispatcher

import (
	"fmt"
	"net"
	"net/http"

	"httpserver/internal/constants"
	"httpserver/internal/utils"
	"httpserver/internal/worker"
)

func Forward(method, path string, conn net.Conn) {
	worker := worker.ChooseWorker()
	if worker == nil {
		utils.WriteHTTPResponse(conn, constants.StatusServiceUnavailable,
			"There are not available workers")
		return
	}

	targetURL := worker.Address + path

	req, err := http.NewRequest(method, targetURL, nil)
	if err != nil {
		utils.WriteHTTPResponse(conn, constants.StatusBadRequest,
			fmt.Sprintf("[Dispatcher] Request build error: %v", err))
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		utils.WriteHTTPResponse(conn, constants.StatusBadGateway,
			fmt.Sprintf("[Dispatcher] Worker error: %v", err))
		return
	}
	defer resp.Body.Close()

	if err := utils.CopyHTTPResponse(conn, resp); err != nil {
		fmt.Printf("[Dispatcher] Copy response error: %v\n", err)
		return
	}

	worker.Completed++
	worker.Load++
}
