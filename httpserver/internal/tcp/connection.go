package tcp

import (
	"bufio"
	"net"

	"httpserver/internal/constants"
	"httpserver/internal/router"
	"httpserver/internal/utils"
)

/**
 * Handles an individual TCP connection from a client.
 *
 * @param {net.Conn} conn - The TCP connection to handle.
 */
func (client *TCPClient) handleConnection(connection net.Conn) {
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

	// Router
	router.HandleRoute(path, connection)
}
