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

func RegisterWorkersFromEnv() {
    workerList := os.Getenv("WORKERS")

    for i, address := range strings.Split(workerList, ",") {
        address = strings.TrimSpace(address)
        if address == "" {
            continue
        }
        fullAddress := "http://" + address
        id := fmt.Sprintf("w%d", i+1)

        RegisterWorker(id, fullAddress)
        log.Printf("[Dispatcher] Worker %s registered", fullAddress)
    }
}

func RegisterWorker(id, address string) {
	workerRegistry.Lock()
	workers[id] = &Worker{ID: id, Address: address, Active: true, LastCheck: time.Now()}
	workerRegistry.Unlock()
}

