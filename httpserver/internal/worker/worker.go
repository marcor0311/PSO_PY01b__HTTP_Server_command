package worker

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

type Worker struct {
	ID        string
	Address   string
	Active    bool
	LastCheck time.Time
	Load      int
	Completed int
}

var (
	workers        = make(map[string]*Worker)
	workerRegistry sync.Mutex
)

var rrIndex int

func RegisterWorkersFromEnv() {
	workerList := os.Getenv("WORKERS")
	seen := make(map[string]bool)

	for i, address := range strings.Split(workerList, ",") {
		address = strings.TrimSpace(address)
		if address == "" {
			continue
		}

		if !strings.Contains(address, ":") {
			address += ":8080"
		}

		if seen[address] {
			continue
		}
		seen[address] = true

		fullAddress := "http://" + address
		id := fmt.Sprintf("w%d", i+1)

		RegisterWorker(id, fullAddress)
	}
}

func RegisterWorker(id, address string) {
	workerRegistry.Lock()
	workers[id] = &Worker{ID: id, Address: address, Active: true, LastCheck: time.Now()}
	workerRegistry.Unlock()
	log.Printf("[Dispatcher] Worker %s registered", address)
}

func ChooseWorker() *Worker {
	workerRegistry.Lock()
	defer workerRegistry.Unlock()

	active := make([]*Worker, 0, len(workers))
	for _, worker := range workers {
		if worker.Active {
			active = append(active, worker)
		}
	}
	if len(active) == 0 {
		return nil
	}

	worker := active[rrIndex%len(active)]
	rrIndex++ 

	return worker
}