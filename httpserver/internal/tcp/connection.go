package tcp

import (
	"bufio"
	"net"
	"strings"

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
		utils.WriteHTTPResponse(connection, "400 Bad Request", "Error reading request line")
		return
	}

	method, path, version, ok := utils.ParseRequestLine(requestLine)
	if !ok {
		utils.WriteHTTPResponse(connection, "400 Bad Request", "Error parsing the request line")
		return
	}

	headers := make(map[string]string)
	for {
		line, err := bufferedReader.ReadString('\n')
		if err != nil || line == "\r\n" {
			break
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			headers[key] = value
		}
	}

	message := HTTPMessage{
		Method:  method,
		Path:    path,
		Version: version,
		Headers: headers,
	}
	client.ReceiveChan <- message

	// Router
	router.HandleRoute(path, connection)
}
