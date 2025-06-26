package router

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"time"

	"httpserver/internal/constants"
	"httpserver/internal/handlers"
	"httpserver/internal/utils"
)

// Fibonacci number handler
func handleFibonacci(conn net.Conn, path string, br *bufio.Reader) {
	defer utils.RecoverAndRespond(conn)
	handlers.TrackWorker("fibonacci", func() {
		query, err := utils.ExtractQuery(path)
		if err != nil {
			utils.WriteHTTPResponse(conn, constants.StatusBadRequest, err.Error())
			return
		}

		numStr := query.Get("num")
		if numStr == "" {
			utils.WriteHTTPResponse(conn, constants.StatusBadRequest, "Missing 'num' query parameter")
			return
		}

		num, err := strconv.Atoi(numStr)
		if err != nil || num <= 0 {
			utils.WriteHTTPResponse(conn, constants.StatusBadRequest, "'num' must be a positive integer")
			return
		}

		result, err := handlers.Fibonacci(num)
		if err != nil {
			utils.WriteHTTPResponse(conn, constants.StatusBadRequest, err.Error())
			return
		}

		response := fmt.Sprintf("The %d Fibonacci number is: %d", num, result)
		utils.WriteHTTPResponse(conn, constants.StatusOK, response)
	})
}

// Random number handler
func handleRandom(conn net.Conn, path string, br *bufio.Reader) {
	defer utils.RecoverAndRespond(conn)
	handlers.TrackWorker("random", func() {
		query, err := utils.ExtractQuery(path)
		if err != nil {
			utils.WriteHTTPResponse(conn, constants.StatusBadRequest, err.Error())
			return
		}

		countStr := query.Get("count")
		minStr := query.Get("min")
		maxStr := query.Get("max")

		count, err1 := strconv.Atoi(countStr)
		min, err2 := strconv.Atoi(minStr)
		max, err3 := strconv.Atoi(maxStr)

		if err1 != nil || err2 != nil || err3 != nil {
			utils.WriteHTTPResponse(conn, constants.StatusBadRequest, "Parameters 'count', 'min', and 'max' must be integers")
			return
		}

		numbers, err := handlers.Random(count, min, max)
		if err != nil {
			utils.WriteHTTPResponse(conn, constants.StatusBadRequest, err.Error())
			return
		}

		response := fmt.Sprintf("Random numbers: %v", numbers)
		utils.WriteHTTPResponse(conn, constants.StatusOK, response)
	})
}

// Reverse string handler
func handleReverse(conn net.Conn, path string, br *bufio.Reader) {
	defer utils.RecoverAndRespond(conn)
	handlers.TrackWorker("reverse", func() {
		query, err := utils.ExtractQuery(path)
		if err != nil {
			utils.WriteHTTPResponse(conn, constants.StatusBadRequest, err.Error())
			return
		}

		text := query.Get("text")
		if text == "" {
			utils.WriteHTTPResponse(conn, constants.StatusBadRequest, "Missing 'text' parameter")
			return
		}

		reversedText := handlers.ReverseString(text)
		utils.WriteHTTPResponse(conn, constants.StatusOK, reversedText)
	})
}

// To uppercase handler
func handleToUpper(conn net.Conn, path string, br *bufio.Reader) {
	defer utils.RecoverAndRespond(conn)
	handlers.TrackWorker("upper", func() {
		query, err := utils.ExtractQuery(path)
		if err != nil {
			utils.WriteHTTPResponse(conn, constants.StatusBadRequest, err.Error())
			return
		}

		text := query.Get("text")
		if text == "" {
			utils.WriteHTTPResponse(conn, constants.StatusBadRequest, "Missing 'text' parameter")
			return
		}

		upperText := handlers.ToUpper(text)
		utils.WriteHTTPResponse(conn, constants.StatusOK, upperText)
	})

}

// Hash converter handler
func handleHash(conn net.Conn, path string, br *bufio.Reader) {
	defer utils.RecoverAndRespond(conn)
	handlers.TrackWorker("upper", func() {})
	query, err := utils.ExtractQuery(path)
	if err != nil {
		utils.WriteHTTPResponse(conn, constants.StatusBadRequest, err.Error())
		return
	}

	text := query.Get("text")
	if text == "" {
		utils.WriteHTTPResponse(conn, constants.StatusBadRequest, "Missing 'text' parameter")
		return
	}

	hash := handlers.HashSHA256(text)
	utils.WriteHTTPResponse(conn, constants.StatusOK, hash)
}

// Time stamp handler
func handleTimestamp(conn net.Conn, path string, br *bufio.Reader) {
	defer utils.RecoverAndRespond(conn)
	handlers.TrackWorker("timestamp", func() {
		timestamp := handlers.Timestamp()
		utils.WriteHTTPResponse(conn, constants.StatusOK, timestamp)
	})

}

// Help handler
func handleHelp(conn net.Conn, path string, br *bufio.Reader) {
	defer utils.RecoverAndRespond(conn)
	handlers.TrackWorker("help", func() {
		help := handlers.HelpText()
		utils.WriteHTTPResponse(conn, constants.StatusOK, help)
	})

}

// Create file handler
func handleCreateFile(conn net.Conn, path string, br *bufio.Reader) {
	defer utils.RecoverAndRespond(conn)
	handlers.TrackWorker("createfile", func() {
		query, err := utils.ExtractQuery(path)
		if err != nil {
			utils.WriteHTTPResponse(conn, constants.StatusBadRequest, err.Error())
			return
		}

		name := query.Get("name")
		content := query.Get("content")
		repeatString := query.Get("repeat")

		if name == "" || content == "" || repeatString == "" {
			utils.WriteHTTPResponse(conn, constants.StatusBadRequest, "Missing 'name', 'content', or 'repeat' parameter")
			return
		}

		repeat, err := strconv.Atoi(repeatString)
		if err != nil || repeat < 1 {
			utils.WriteHTTPResponse(conn, constants.StatusBadRequest, "'repeat' has to be a positive integer")
			return
		}

		err = handlers.CreateFile(name, content, repeat)
		if err != nil {
			utils.WriteHTTPResponse(conn, "500 Internal Server Error", err.Error())
			return
		}

		utils.WriteHTTPResponse(conn, constants.StatusOK, fmt.Sprintf("File '%s' created successfully", name))
	})

}

// Delete file handler
func handleDeleteFile(conn net.Conn, path string, br *bufio.Reader) {
	defer utils.RecoverAndRespond(conn)
	handlers.TrackWorker("deletefile", func() {
		query, err := utils.ExtractQuery(path)
		if err != nil {
			utils.WriteHTTPResponse(conn, constants.StatusBadRequest, err.Error())
			return
		}

		name := query.Get("name")
		if name == "" {
			utils.WriteHTTPResponse(conn, constants.StatusBadRequest, "Missing 'name' parameter")
			return
		}

		err = handlers.DeleteFile(name)
		if err != nil {
			utils.WriteHTTPResponse(conn, "500 Internal Server Error", err.Error())
			return
		}

		utils.WriteHTTPResponse(conn, constants.StatusOK, fmt.Sprintf("File '%s' deleted successfully", name))
	})

}

// Simulate task handler
func handleSimulate(conn net.Conn, path string, br *bufio.Reader) {
	defer utils.RecoverAndRespond(conn)
	handlers.TrackWorker("simulate", func() {
		query, err := utils.ExtractQuery(path)
		if err != nil {
			utils.WriteHTTPResponse(conn, constants.StatusBadRequest, err.Error())
			return
		}

		secondsStr := query.Get("seconds")
		task := query.Get("task")

		if secondsStr == "" || task == "" {
			utils.WriteHTTPResponse(conn, constants.StatusBadRequest, "Missing 'seconds' or 'task' parameter")
			return
		}

		seconds, err := strconv.Atoi(secondsStr)
		if err != nil || seconds < 1 {
			utils.WriteHTTPResponse(conn, constants.StatusBadRequest, "'seconds' must be a positive integer")
			return
		}

		message := handlers.SimulateTask(seconds, task)
		utils.WriteHTTPResponse(conn, constants.StatusOK, message)
	})

}

// Sleep handler
func handleSleep(conn net.Conn, path string, br *bufio.Reader) {
	defer utils.RecoverAndRespond(conn)
	handlers.TrackWorker("sleep", func() {
		query, err := utils.ExtractQuery(path)
		if err != nil {
			utils.WriteHTTPResponse(conn, constants.StatusBadRequest, err.Error())
			return
		}

		secondsStr := query.Get("seconds")
		if secondsStr == "" {
			utils.WriteHTTPResponse(conn, constants.StatusBadRequest, "Missing 'seconds' parameter")
			return
		}

		seconds, err := strconv.Atoi(secondsStr)
		if err != nil {
			utils.WriteHTTPResponse(conn, constants.StatusBadRequest, "'seconds' must be a positive integer")
			return
		}

		message := handlers.Sleep(seconds)
		utils.WriteHTTPResponse(conn, constants.StatusOK, message)
	})

}

// /loadtest?tasks=n&sleep=s
func handleLoadTest(conn net.Conn, path string, br *bufio.Reader) {
	defer utils.RecoverAndRespond(conn)
	handlers.TrackWorker("loadtest", func() {
		query, err := utils.ExtractQuery(path)
		if err != nil {
			utils.WriteHTTPResponse(conn, constants.StatusBadRequest, err.Error())
			return
		}

		tasksStr := query.Get("tasks")
		sleepStr := query.Get("sleep")

		if tasksStr == "" || sleepStr == "" {
			utils.WriteHTTPResponse(conn, constants.StatusBadRequest, "Missing 'tasks' or 'sleep' parameter")
			return
		}

		tasks, err := strconv.Atoi(tasksStr)
		if err != nil || tasks <= 0 {
			utils.WriteHTTPResponse(conn, constants.StatusBadRequest, "'tasks' must be a positive integer")
			return
		}

		sleepSec, err := strconv.Atoi(sleepStr)
		if err != nil || sleepSec < 0 {
			utils.WriteHTTPResponse(conn, constants.StatusBadRequest, "'sleep' must be a non-negative integer")
			return
		}

		duration, err := handlers.SimulateLoad(tasks, sleepSec)
		if err != nil {
			utils.WriteHTTPResponse(conn, constants.StatusInternalServerError, err.Error())
			return
		}

		message := "Simulated " + strconv.Itoa(tasks) + " tasks with " + strconv.Itoa(sleepSec) + "s sleep each.\nTotal time: " + duration.String()
		utils.WriteHTTPResponse(conn, constants.StatusOK, message)
	})

}

// Status handler
func handleStatus(conn net.Conn, path string, br *bufio.Reader) {
	defer utils.RecoverAndRespond(conn)
	handlers.TrackWorker("status", func() {
		statusJSON, err := handlers.GetStatusJSON()
		if err != nil {
			utils.WriteHTTPResponse(conn, constants.StatusInternalServerError, "Failed to generate status report")
			return
		}

		utils.WriteHTTPResponse(conn, constants.StatusOK, statusJSON)
	})
}

func handleWordCountChunk(conn net.Conn, path string, br *bufio.Reader) {
	defer utils.RecoverAndRespond(conn)

	// discard headers (read until blank line)
	for {
		line, err := br.ReadString('\n')
		if err != nil {
			utils.WriteHTTPResponse(conn, constants.StatusBadRequest, "cannot read headers")
			return
		}
		if line == "\r\n" || line == "\n" {
			break
		}
	}

	var req struct {
		ID    int    `json:"id"`
		Chunk string `json:"chunk"`
	}
	if err := json.NewDecoder(br).Decode(&req); err != nil {
		utils.WriteHTTPResponse(conn, constants.StatusBadRequest, "invalid JSON")
		return
	}

	freq := handlers.CountWords(req.Chunk)

	resBytes, _ := json.Marshal(struct {
		ID   int            `json:"id"`
		Freq map[string]int `json:"freq"`
	}{req.ID, freq})

	utils.WriteHTTPResponse(conn, constants.StatusOK, string(resBytes))
}

// Handle Monte Carlo simulation
func handleMontecarlo(conn net.Conn, path string, br *bufio.Reader) {
	defer utils.RecoverAndRespond(conn)
	handlers.TrackWorker("montecarlo", func() {
		for {
			line, err := br.ReadString('\n')
			if err != nil {
				utils.WriteHTTPResponse(conn, constants.StatusBadRequest, "cannot read headers")
				return
			}
			if line == "\r\n" || line == "\n" {
				break
			}
		}

		var req struct {
			ID     int `json:"id"`
			Points int `json:"points"`
		}
		if err := json.NewDecoder(br).Decode(&req); err != nil {
			utils.WriteHTTPResponse(conn, constants.StatusBadRequest, "invalid JSON")
			return
		}

		rand.Seed(time.Now().UnixNano())
		inside := 0
		for i := 0; i < req.Points; i++ {
			x := rand.Float64()
			y := rand.Float64()
			if x*x+y*y <= 1.0 {
				inside++
			}
		}

		resBytes, _ := json.Marshal(struct {
			ID     int `json:"id"`
			Inside int `json:"inside"`
		}{ID: req.ID, Inside: inside})

		utils.WriteHTTPResponse(conn, constants.StatusOK, string(resBytes))
	})
}

// Not found handler
func handleNotFound(conn net.Conn, path string, _ *bufio.Reader) {
	utils.WriteHTTPResponse(conn, constants.StatusNotFound, fmt.Sprintf("Unknown path: %s", path))
}
