package tcp

import (
	"fmt"
	"io"
	"net"
	"strings"
)

type HTTPMessage struct {
	Method  string
	Path    string
	Version string
	Headers map[string]string
	Body    string
}

/**
 * Parses the HTTP request line into method, path, and version.
 * Expected format: "GET /path HTTP/1.0"
 *
 * @param {string} line - The request line from the HTTP client.
 * @returns {string, string, string, bool} - Method, path, version, and success flag.
 */
func parseRequestLine(line string) (method, path, version string, ok bool) {
	parts := strings.Fields(line)
	if len(parts) != 3 {
		return
	}
	return parts[0], parts[1], parts[2], true
}

/**
 * Writes an HTTP/1.0 response to the TCP connection.
 * Includes status, headers and body.
 *
 * @param {net.Conn} conn - The TCP connection to write the response to.
 * @param {string} status - HTTP status line (e.g., "200 OK").
 * @param {string} body - The plain text body to send in the response.
 */
func writeHTTPResponse(connection net.Conn, status string, body string) {
	response := fmt.Sprintf(
		"HTTP/1.0 %s\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s",
		status,
		len(body),
		body,
	)
	io.WriteString(connection, response)
}
