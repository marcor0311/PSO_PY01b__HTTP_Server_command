package tcp

import (
	"bufio"
	"net"

	"httpserver/internal/constants"
	"httpserver/internal/dispatcher"
	"httpserver/internal/router"
	"httpserver/internal/utils"
)

/**
 * Handles an individual TCP connection from a client.
 *
 * @param {net.Conn} conn - The TCP connection to handle.
 */
func (client *TCPClient) handleWorkerConnection(connection net.Conn) {
	defer connection.Close()
	bufferedReader := bufio.NewReader(connection)

	requestLine, err := bufferedReader.ReadString('\n')
	if err != nil {
		utils.WriteHTTPResponse(connection, constants.StatusBadRequest, "Error reading request line")
		return
	}

	method, path, version, ok := utils.ParseRequestLine(requestLine)
	if !ok {
		utils.WriteHTTPResponse(connection, constants.StatusBadRequest, "Error parsing the request line")
		return
	}

	message := HTTPMessage{
		Method:  method,
		Path:    path,
		Version: version,
	}
	client.ReceiveChan <- message

	// 👉 router now receives the same bufferedReader positioned at headers
	router.HandleRoute(path, connection, bufferedReader)
}

/**
 * Handles an incoming client connection on the Dispatcher.
 * Reads the request line, decides whether the route requires parallel processing
 * or simple forwarding.
 *
 * @param {net.Conn} conn - The TCP connection to handle.
 */
func (client *TCPClient) handleDispatcherConnection(connection net.Conn) {
	defer connection.Close()
	bufferedReader := bufio.NewReader(connection)

	reqLine, err := bufferedReader.ReadString('\n')
	if err != nil {
		utils.WriteHTTPResponse(connection, constants.StatusBadRequest, "Error reading request line")
		return
	}

	method, path, _, ok := utils.ParseRequestLine(reqLine)
	if !ok {
		utils.WriteHTTPResponse(connection, constants.StatusBadRequest, "Error parsing the request line")
		return
	}

	handled := router.HandleDispatcherRouter(path, connection, bufferedReader)
	if handled {
		return
	}

	dispatcher.Forward(method, path, connection)
}
