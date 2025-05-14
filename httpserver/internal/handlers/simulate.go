package handlers

import (
	"fmt"
	"time"
)

// /simulate?seconds=s&task=name: Simulates a task that takes 's' seconds to complete.
func SimulateTask(seconds int, taskName string) string {
	if seconds <= 0 {
		return "Error: Time must be a positive number"
	}
	time.Sleep(time.Duration(seconds) * time.Second)
	return fmt.Sprintf("Task '%s' completed after %d seconds", taskName, seconds)
}

// /sleep?seconds=s: Simulates latency.
func Sleep(seconds int) string {
	if seconds <= 0 {
		return "Error: Time must be a positive number"
	}
	time.Sleep(time.Duration(seconds) * time.Second)
	return fmt.Sprintf("Slept for %d seconds", seconds)
}