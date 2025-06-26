package handlers

import (
	"encoding/json"
	"fmt"
	"httpserver/internal/worker"
	"os"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

type WorkerStatus struct {
	PID    int    `json:"pid"`
	Task   string `json:"task"`
	Status string `json:"status"`
}

type ServerStatus struct {
	Uptime        string         `json:"uptime"`
	TotalRequests int64          `json:"total_requests"`
	MainPID       int            `json:"main_pid"`
	Workers       []WorkerStatus `json:"workers"`
}

var startTime = time.Now()
var requestCount int64
var workerLock sync.RWMutex
var workers = make(map[int]WorkerStatus)
var workerCounter int64

func incrementRequestCount() {
	atomic.AddInt64(&requestCount, 1)
}

func RegisterWorker(id int64, task string) {
	workerLock.Lock()
	defer workerLock.Unlock()
	workers[int(id)] = WorkerStatus{
		PID:    int(id),
		Task:   task,
		Status: "busy",
	}
}

func SetWorkerAvailable(id int64) {
	workerLock.Lock()
	defer workerLock.Unlock()
	worker, exists := workers[int(id)]
	if exists {
		worker.Status = "available"
		workers[int(id)] = worker
	}
}

func TrackWorker(taskName string, fn func()) {
	id := atomic.AddInt64(&workerCounter, 1)
	incrementRequestCount()
	RegisterWorker(id, taskName)
	fn()
	SetWorkerAvailable(id)
}

func GetStatusJSON() (string, error) {
	uptime := time.Since(startTime).Truncate(time.Second)

	workerLock.RLock()
	defer workerLock.RUnlock()

	activeWorkers := make([]WorkerStatus, 0, len(workers))
	for _, w := range workers {
		activeWorkers = append(activeWorkers, w)
	}

	status := ServerStatus{
		Uptime:        uptime.String(),
		TotalRequests: atomic.LoadInt64(&requestCount),
		MainPID:       os.Getpid(),
		Workers:       activeWorkers,
	}

	jsonBytes, err := json.MarshalIndent(status, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal status: %w", err)
	}
	return string(jsonBytes), nil
}

func GetWorkerInformation() (string, error) {
	worker.WorkerRegistry.Lock()
	list := make([]worker.Worker, 0, len(worker.Workers))
	for _, w := range worker.Workers {
		list = append(list, *w)
	}
	worker.WorkerRegistry.Unlock()

	sort.Slice(list, func(i, j int) bool { return list[i].ID < list[j].ID })

	for i := range list {
		list[i].LastCheck = list[i].LastCheck.In(time.UTC)
	}

	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return "", fmt.Errorf("Failed parse response")
	}
	return string(data), nil
}
