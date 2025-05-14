package router

import (
	"fmt"
	"net"
	"strconv"

	"httpserver/internal/handlers"
	"httpserver/internal/utils"

)

// Fibonacci number handler
func handleFibonacci(conn net.Conn, path string) {
	defer utils.RecoverAndRespond(conn)

    query, err := utils.ExtractQuery(path)
    if err != nil {
        utils.WriteHTTPResponse(conn, "400 Bad Request", err.Error())
        return
    }

    numStr := query.Get("num")
    if numStr == "" {
        utils.WriteHTTPResponse(conn, "400 Bad Request", "Missing 'num' query parameter")
        return
    }

    num, err := strconv.Atoi(numStr)
    if err != nil || num <= 0 {
        utils.WriteHTTPResponse(conn, "400 Bad Request", "'num' must be a positive integer")
        return
    }

    result, err := handlers.Fibonacci(num)
    if err != nil {
        utils.WriteHTTPResponse(conn, "400 Bad Request", err.Error())
        return
    }

    response := fmt.Sprintf("The %d Fibonacci number is: %d", num, result)
    utils.WriteHTTPResponse(conn, "200 OK", response)
}

// Random number handler
func handleRandom(conn net.Conn, path string) {
	defer utils.RecoverAndRespond(conn)

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

// Reverse string handler
func handleReverse(conn net.Conn, path string) {
	defer utils.RecoverAndRespond(conn)

	query, err := utils.ExtractQuery(path)
	if err != nil {
		utils.WriteHTTPResponse(conn, "400 Bad Request", err.Error())
		return
	}

	text := query.Get("text")
	if text == "" {
		utils.WriteHTTPResponse(conn, "400 Bad Request", "Missing 'text' parameter")
		return
	}

	reversedText := handlers.ReverseString(text)
	utils.WriteHTTPResponse(conn, "200 OK", reversedText)
}

// To uppercase handler
func handleToUpper(conn net.Conn, path string) {
	defer utils.RecoverAndRespond(conn)

	query, err := utils.ExtractQuery(path)
	if err != nil {
		utils.WriteHTTPResponse(conn, "400 Bad Request", err.Error())
		return
	}

	text := query.Get("text")
	if text == "" {
		utils.WriteHTTPResponse(conn, "400 Bad Request", "Missing 'text' parameter")
		return
	}

	upperText := handlers.ToUpper(text)
	utils.WriteHTTPResponse(conn, "200 OK", upperText)
}

// Hash converter handler 
func handleHash(conn net.Conn, path string) {
	defer utils.RecoverAndRespond(conn)

	query, err := utils.ExtractQuery(path)
	if err != nil {
		utils.WriteHTTPResponse(conn, "400 Bad Request", err.Error())
		return
	}

	text := query.Get("text")
	if text == "" {
		utils.WriteHTTPResponse(conn, "400 Bad Request", "Missing 'text' parameter")
		return
	}

	hash := handlers.HashSHA256(text)
	utils.WriteHTTPResponse(conn, "200 OK", hash)
}

// Time stamp handler
func handleTimestamp(conn net.Conn, path string) {
	defer utils.RecoverAndRespond(conn)

	timestamp := handlers.Timestamp()
	utils.WriteHTTPResponse(conn, "200 OK", timestamp)
}

// Help handler
func handleHelp(conn net.Conn, path string) {
	defer utils.RecoverAndRespond(conn)

	help := handlers.HelpText()
	utils.WriteHTTPResponse(conn, "200 OK", help)
}


func handleCreateFile(conn net.Conn, path string) {
	query, err := utils.ExtractQuery(path)
	if err != nil {
		utils.WriteHTTPResponse(conn, "400 Bad Request", err.Error())
		return
	}

	name := query.Get("name")
	content := query.Get("content")
	repeatStr := query.Get("repeat")

	if name == "" || content == "" || repeatStr == "" {
		utils.WriteHTTPResponse(conn, "400 Bad Request", "Missing 'name', 'content', or 'repeat' parameter")
		return
	}

	repeat, err := strconv.Atoi(repeatStr)
	if err != nil || repeat < 1 {
		utils.WriteHTTPResponse(conn, "400 Bad Request", "'repeat' has to be a positive integer")
		return
	}

	err = handlers.CreateFile(name, content, repeat)
	if err != nil {
		utils.WriteHTTPResponse(conn, "500 Internal Server Error", err.Error())
		return
	}

	utils.WriteHTTPResponse(conn, "200 OK", fmt.Sprintf("File '%s' created successfully", name))
}

func handleDeleteFile(conn net.Conn, path string) {
	query, err := utils.ExtractQuery(path)
	if err != nil {
		utils.WriteHTTPResponse(conn, "400 Bad Request", err.Error())
		return
	}

	name := query.Get("name")
	if name == "" {
		utils.WriteHTTPResponse(conn, "400 Bad Request", "Missing 'name' parameter")
		return
	}

	err = handlers.DeleteFile(name)
	if err != nil {
		utils.WriteHTTPResponse(conn, "500 Internal Server Error", err.Error())
		return
	}

	utils.WriteHTTPResponse(conn, "200 OK", fmt.Sprintf("File '%s' deleted successfully", name))
}

func handleSimulate(conn net.Conn, path string) {
	utils.WriteHTTPResponse(conn, "200 OK", "[TODO] Simulate Task")
}
func handleSleep(conn net.Conn, path string) { utils.WriteHTTPResponse(conn, "200 OK", "[TODO] Sleep") }
func handleLoadTest(conn net.Conn, path string) {
	utils.WriteHTTPResponse(conn, "200 OK", "[TODO] Load Test")
}
func handleStatus(conn net.Conn, path string) {
	utils.WriteHTTPResponse(conn, "200 OK", "Server is running")
}
func handleNotFound(conn net.Conn, path string) {
	utils.WriteHTTPResponse(conn, "404 Not Found", fmt.Sprintf("Unknown path: %s", path))
}
