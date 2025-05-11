package tcp

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

/**
 * Handles an individual TCP connection from a client.
 *
 * @param {net.Conn} conn - The TCP connection to handle.
 */
func (client *TCPClient) handleConnection(connection net.Conn) {
	defer connection.Close()
	bufferedReader := bufio.NewReader(connection)

	// Reads the HTTP request line
	requestLine, err := bufferedReader.ReadString('\n')
	if err != nil {
		writeHTTPResponse(connection, "400 Bad Request", "Error reading request line")
		return
	}

	// Parse method, path, and HTTP version
	method, path, version, ok := parseRequestLine(requestLine)
	if !ok {
		writeHTTPResponse(connection, "400 Bad Request", " Error parsing the request line")
		return
	}

	// Parse headers
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

	// Creates the parsed HTTP request
	message := HTTPMessage{
		Method:  method,
		Path:    path,
		Version: version,
		Headers: headers,
	}
	client.ReceiveChan <- message

	var responseStatus string
	var responseBody string

	switch {
	case path == "/status":
		responseStatus = "200 OK"
		responseBody = "Server is running"
	default:
		responseStatus = "404 Not Found"
		responseBody = fmt.Sprintf("Unknown path: %s", path)
	}

	writeHTTPResponse(connection, responseStatus, responseBody)
}