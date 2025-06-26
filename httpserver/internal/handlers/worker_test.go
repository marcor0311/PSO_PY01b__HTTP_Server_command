package handlers

import (
	"encoding/json"
	"httpserver/internal/worker"
	"os"
	"strings"
	"sync"
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

func TestGetWorkerInformation(t *testing.T) {
	// Backup original state
	originalWorkers := worker.Workers

	// Setup mock workers
	mockTime := time.Date(2025, 6, 26, 10, 0, 0, 0, time.UTC)
	worker.Workers = map[string]*worker.Worker{
		"w2": {ID: "w2", LastCheck: mockTime},
		"w1": {ID: "w1", LastCheck: mockTime.Add(-1 * time.Minute)},
	}
	worker.WorkerRegistry = sync.Mutex{}

	expected := []worker.Worker{
		{ID: "w1", LastCheck: mockTime.Add(-1 * time.Minute).In(time.UTC)},
		{ID: "w2", LastCheck: mockTime.In(time.UTC)},
	}

	t.Run("returns sorted and JSON-formatted workers", func(t *testing.T) {
		result, err := GetWorkerInformation()
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		var parsed []worker.Worker
		if err := json.Unmarshal([]byte(result), &parsed); err != nil {
			t.Fatalf("Failed to parse JSON: %v", err)
		}

		if len(parsed) != len(expected) {
			t.Errorf("Expected %d workers, got %d", len(expected), len(parsed))
		}

		for i := range expected {
			if parsed[i].ID != expected[i].ID || !parsed[i].LastCheck.Equal(expected[i].LastCheck) {
				t.Errorf("Mismatch at index %d: got %+v, expected %+v", i, parsed[i], expected[i])
			}
		}
	})

	// Restore original state
	worker.Workers = originalWorkers
}
