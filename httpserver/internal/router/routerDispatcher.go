package router

import (
	"bufio"
	"net"
	"net/url"

	"httpserver/internal/constants"
	"httpserver/internal/handlers"
	"httpserver/internal/utils"
)

// Ping
func handlePing(conn net.Conn, path string, br *bufio.Reader) {
	utils.WriteHTTPResponse(conn, constants.StatusOK, "pong")
}

// Workers
func handleWorkers(conn net.Conn, path string, br *bufio.Reader) {
	defer utils.RecoverAndRespond(conn)
	handlers.TrackWorker("workers", func() {
		statusJSON, err := handlers.GetWorkerInformation()
		if err != nil {
			utils.WriteHTTPResponse(conn, constants.StatusInternalServerError, "Failed to generate worker report")
			return
		}

		utils.WriteHTTPResponse(conn, constants.StatusOK, statusJSON)
	})
}

func handlParallelWordCount(conn net.Conn, path string, br *bufio.Reader) {
	defer utils.RecoverAndRespond(conn)

	handlers.TrackWorker("wordcount", func() {
		url, err := url.Parse(path)
		if err != nil {
			utils.WriteHTTPResponse(conn, constants.StatusBadRequest, "Malformed URL")
			return
		}
		fileURL := url.Query().Get("url")
		if fileURL == "" {
			utils.WriteHTTPResponse(conn, constants.StatusBadRequest, "URL parameter missing")
			return
		}

		resultJSON, err := HandleParallelWordCount(fileURL)
		if err != nil {
			utils.WriteHTTPResponse(conn, constants.StatusBadGateway, err.Error())
			return
		}

		utils.WriteHTTPResponse(conn, constants.StatusOK, resultJSON)
	})
}
