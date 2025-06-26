package handlers

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
	"time"
)

func resetMetricsState() {
	startTime = time.Now()
	requestCount = 0
	workerCounter = 0
	workerLock.Lock()
	defer workerLock.Unlock()
	workers = make(map[int]WorkerStatus)
}

func TestRegisterAndTrackWorker(t *testing.T) {
	resetMetricsState()

	taskName := "test-task"
	TrackWorker(taskName, func() {})

	jsonStr, err := GetStatusJSON()
	if err != nil {
		t.Fatalf("GetStatusJSON returned error: %v", err)
	}

	if !strings.Contains(jsonStr, taskName) {
		t.Errorf("Expected task name %q in JSON, got: %s", taskName, jsonStr)
	}

	var status ServerStatus
	if err := json.Unmarshal([]byte(jsonStr), &status); err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	if status.TotalRequests != 1 {
		t.Errorf("Expected TotalRequests to be 1, got %d", status.TotalRequests)
	}

	if len(status.Workers) != 1 {
		t.Fatalf("Expected 1 worker, got %d", len(status.Workers))
	}

	worker := status.Workers[0]
	if worker.Status != "available" {
		t.Errorf("Expected worker status 'available', got %q", worker.Status)
	}

	if worker.Task != taskName {
		t.Errorf("Expected task name %q, got %q", taskName, worker.Task)
	}

	if status.MainPID != os.Getpid() {
		t.Errorf("Expected PID %d, got %d", os.Getpid(), status.MainPID)
	}
}
