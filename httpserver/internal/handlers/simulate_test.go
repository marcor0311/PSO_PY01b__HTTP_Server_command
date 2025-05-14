package handlers_test

import (
	"strings"
	"testing"
	"time"

	"httpserver/internal/handlers"
)

// Simulate 
func TestSimulateTask(t *testing.T) {
	taskName := "simulated_task"
	seconds := 5

	start := time.Now()
	message := handlers.SimulateTask(seconds, taskName)
	duration := time.Since(start)

	if duration < time.Duration(seconds)*time.Second {
		t.Errorf("Simulated task took less time than expected. Expected: %d seconds, got: %v", seconds, duration)
	}

	if !strings.Contains(message, taskName) {
		t.Errorf("Message does not contain task name. Expected: %s, got: %s", taskName, message)
	}
}

func TestSimulateTask_NegativeSeconds(t *testing.T) {
	taskName := "simulated_task"
	seconds := -5

	message := handlers.SimulateTask(seconds, taskName)

	if message != "Error: Time must be a positive number" {
		t.Errorf("Expected error message for negative seconds, got: %s", message)
	}
}

func TestSimulateTask_ZeroSeconds(t *testing.T) {
	taskName := "simulated_task"
	seconds := 0

	message := handlers.SimulateTask(seconds, taskName)

	if message != "Error: Time must be a positive number" {
		t.Errorf("Expected error message for zero seconds, got: %s", message)
	}
}

// Sleep

func TestSleep(t *testing.T) {
	seconds := 3

	start := time.Now()
	message := handlers.Sleep(seconds)
	duration := time.Since(start)

	if duration < time.Duration(seconds)*time.Second {
		t.Errorf("Sleep function took less time than expected. Expected: %d seconds, got: %v", seconds, duration)
	}

	if !strings.Contains(message, "Slept for") {
		t.Errorf("Message does not contain sleep duration. Expected: 'Slept for', got: %s", message)
	}
}

func TestSleep_NegativeSeconds(t *testing.T) {
	seconds := -3

	message := handlers.Sleep(seconds)

	if message != "Error: Time must be a positive number" {
		t.Errorf("Expected error message for negative seconds, got: %s", message)
	}
}