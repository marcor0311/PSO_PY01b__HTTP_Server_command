package router

import (
	"fmt"
	"net"
	"strconv"

	"httpserver/internal/handlers"
	"httpserver/internal/utils"

)

func handleFibonacci(conn net.Conn, path string) {
    // Extract query parameters from the path
    query, err := utils.ExtractQuery(path)
    if err != nil {
        utils.WriteHTTPResponse(conn, "400 Bad Request", err.Error())
        return
    }

    // Get the "num" parameter
    numStr := query.Get("num")
    if numStr == "" {
        utils.WriteHTTPResponse(conn, "400 Bad Request", "Missing 'num' query parameter")
        return
    }

    // Convert "num" to an integer
    num, err := strconv.Atoi(numStr)
    if err != nil || num <= 0 {
        utils.WriteHTTPResponse(conn, "400 Bad Request", "'num' must be a positive integer")
        return
    }

    // Calculate the Nth Fibonacci number
    result, err := handlers.Fibonacci(num)
    if err != nil {
        utils.WriteHTTPResponse(conn, "400 Bad Request", err.Error())
        return
    }

    // Format the response
    response := fmt.Sprintf("The %d Fibonacci number is: %d", num, result)
    utils.WriteHTTPResponse(conn, "200 OK", response)
}

func handleCreateFile(conn net.Conn, path string) {
	utils.WriteHTTPResponse(conn, "200 OK", "[TODO] Create File")
}
func handleDeleteFile(conn net.Conn, path string) {
	utils.WriteHTTPResponse(conn, "200 OK", "[TODO] Delete File")
}
func handleReverse(conn net.Conn, path string) {
	utils.WriteHTTPResponse(conn, "200 OK", "[TODO] Reverse Text")
}
func handleToUpper(conn net.Conn, path string) {
	utils.WriteHTTPResponse(conn, "200 OK", "[TODO] To Upper")
}

// Random number handler
func handleRandom(conn net.Conn, path string) {
	query, err := utils.ExtractQuery(path)
	if err != nil {
		utils.WriteHTTPResponse(conn, "400 Bad Request", err.Error())
		return
	}

	countStr := query.Get("count")
	minStr := query.Get("min")
	maxStr := query.Get("max")

	count, err1 := strconv.Atoi(countStr)
	min, err2 := strconv.Atoi(minStr)
	max, err3 := strconv.Atoi(maxStr)

	if err1 != nil || err2 != nil || err3 != nil {
		utils.WriteHTTPResponse(conn, "400 Bad Request", "Parameters 'count', 'min', and 'max' must be integers")
		return
	}

	numbers, err := handlers.Random(count, min, max)
	if err != nil {
		utils.WriteHTTPResponse(conn, "400 Bad Request", err.Error())
		return
	}

	response := fmt.Sprintf("Random numbers: %v", numbers)
	utils.WriteHTTPResponse(conn, "200 OK", response)
}

func handleTimestamp(conn net.Conn, path string) {
	utils.WriteHTTPResponse(conn, "200 OK", "[TODO] Timestamp")
}
func handleHash(conn net.Conn, path string) { utils.WriteHTTPResponse(conn, "200 OK", "[TODO] Hash") }
func handleSimulate(conn net.Conn, path string) {
	utils.WriteHTTPResponse(conn, "200 OK", "[TODO] Simulate Task")
}
func handleSleep(conn net.Conn, path string) { utils.WriteHTTPResponse(conn, "200 OK", "[TODO] Sleep") }
func handleLoadTest(conn net.Conn, path string) {
	utils.WriteHTTPResponse(conn, "200 OK", "[TODO] Load Test")
}
func handleHelp(conn net.Conn, path string) { utils.WriteHTTPResponse(conn, "200 OK", "[TODO] Help") }
func handleStatus(conn net.Conn, path string) {
	utils.WriteHTTPResponse(conn, "200 OK", "Server is running")
}
func handleNotFound(conn net.Conn, path string) {
	utils.WriteHTTPResponse(conn, "404 Not Found", fmt.Sprintf("Unknown path: %s", path))
}
