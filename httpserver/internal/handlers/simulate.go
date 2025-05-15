package handlers

import (
	"fmt"
	"time"
	"sync"
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

// /loadtest?tasks=n&sleep=x: SimulateLoad runs 'tasks' goroutines, each sleeping for 'n' seconds.
func SimulateLoad(tasks, sleepSec int) (time.Duration, error) {
	if tasks <= 0 || sleepSec < 0 {
		return 0, fmt.Errorf("invalid input: tasks must be > 0 and sleep >= 0")
	}

	var waitGroup sync.WaitGroup
	start := time.Now()

	for i := 0; i < tasks; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			time.Sleep(time.Duration(sleepSec) * time.Second)
		}()
	}

	waitGroup.Wait()
	return time.Since(start), nil
}