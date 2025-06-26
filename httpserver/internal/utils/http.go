package utils

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"httpserver/internal/constants"
)

/**
 * Parses the HTTP request line into method, path, and version.
 * Expected format: "GET /path HTTP/1.0"
 *
 * @param {string} line - The request line from the HTTP client.
 * @returns {string, string, string, bool} - Method, path, version, and success flag.
 */
func ParseRequestLine(line string) (method, path, version string, ok bool) {
	fields := strings.Fields(line)
	if len(fields) != 3 {
		return "", "", "", false
	}
	method, path, version = fields[0], fields[1], fields[2]
	return method, path, version, true
}

/**
 * Writes an HTTP/1.0 response to the TCP connection.
 * Includes status, headers and body.
 *
 * @param {net.Conn} conn - The TCP connection to write the response to.
 * @param {string} status - HTTP status line (e.g., "constants.StatusOK").
 * @param {string} body - The plain text body to send in the response.
 */
func WriteHTTPResponse(conn net.Conn, status, body string) {
	contentType := "text/plain"
	bodyBytes := []byte(body)
	contentLength := len(bodyBytes)

	headers := fmt.Sprintf("HTTP/1.0 %s\r\nContent-Type: %s\r\nContent-Length: %d\r\n\r\n",
		status, contentType, contentLength)

	conn.Write([]byte(headers))
	conn.Write(bodyBytes)
}

/**
 * Copy an http response back to the client over the raw TCP connection.
 * Writes the status line, all headers, a blank line, and finally the response body.
 * @param {net.Conn} destination - Client TCP connection to write to.
 * @param {*http.Response} response - Response received from the worker.
 */
func CopyHTTPResponse(destination net.Conn, response *http.Response) error {
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// Re-use the helper that already writes status line, headers and body.
	WriteHTTPResponse(destination, response.Status, string(bodyBytes))
	return nil
}

/**
 * ExtractQuery parses and returns the query parameters from the URL path.
 */
func ExtractQuery(path string) (url.Values, error) {
	parts := strings.SplitN(path, "?", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("missing query parameters in path: %q", path)
	}

	query, err := url.ParseQuery(parts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid query format: %v", err)
	}

	return query, nil
}

func ReadRequestBody(br *bufio.Reader, headers map[string]string) ([]byte, error) {
	clen, err := strconv.Atoi(headers["Content-Length"])
	if err != nil {
		return nil, fmt.Errorf("Missing or invalid Content-Length")
	}
	buf := make([]byte, clen)
	_, err = io.ReadFull(br, buf)
	return buf, err
}

/**
 * RecoverAndRespond catches a system error and sends a 500 response.
 */
func RecoverAndRespond(conn net.Conn) {
	if r := recover(); r != nil {
		WriteHTTPResponse(conn, constants.StatusInternalServerError, fmt.Sprintf("Internal server error: %v", r))
	}
}
