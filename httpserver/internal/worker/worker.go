package worker

import (
	"fmt"
	"httpserver/internal/constants"
	"log"
	"net/http"
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
	Workers        = make(map[string]*Worker)
	WorkerRegistry sync.Mutex
)

var rrIndex int

/**
 * Reads the WORKERS environment variable and registers every worker.
 */
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

/**
 * Adds workers entries to the registry, marks it active, and logs the event.
 *
 * @param {string} id - Logical identifier for the worker.
 * @param {string} address - Base URL where the worker is reachable.
 */
func RegisterWorker(id, address string) {
	WorkerRegistry.Lock()
	Workers[id] = &Worker{ID: id, Address: address, Active: true, LastCheck: time.Now()}
	WorkerRegistry.Unlock()
	log.Printf("[Dispatcher] Worker %s registered", address)
}

/**
 * Selects one active worker using round-robin logic.
 *
 * @returns {*Worker} - Pointer to the chosen worker or nil.
 */
func ChooseWorker() *Worker {
	WorkerRegistry.Lock()
	defer WorkerRegistry.Unlock()

	active := []*Worker{}
	for _, w := range Workers {
		if w.Active {
			active = append(active, w)
		}
	}
	if len(active) == 0 {
		return nil
	}

	selected := active[rrIndex%len(active)]
	rrIndex++
	return selected
}

/**
 * Marks worker as inactive
 */
func (worker *Worker) MarkInactive() {
	WorkerRegistry.Lock()
	defer WorkerRegistry.Unlock()
	worker.Active = false
	log.Printf("[Dispatcher] Marked worker %s as inactive", worker.Address)
}

/**
 * Sends HTTP request to the worker
 *
 * @param {string} endpoint - Enpoint.
 * @param {body} string - Request body.
 */
func SendRequestToWorker(endpoint string, body string) (*http.Response, *Worker, error) {
	worker := ChooseWorker()
	if worker == nil {
		return nil, nil, fmt.Errorf("There are no active workers")
	}

	log.Printf("[Dispatcher] Sending request (%s) to worker %s", endpoint, worker.Address)

	url := worker.Address + endpoint
	req, _ := http.NewRequest("POST", url, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("[Dispatcher] Error contacting %s, marking as inactive", worker.Address)
		worker.MarkInactive()
		return nil, worker, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("[Dispatcher] Worker %s returned %s, marking as inactive", worker.Address, resp.Status)
		resp.Body.Close()
		worker.MarkInactive()
		return nil, worker, fmt.Errorf("Bad response from %s: %s", worker.Address, resp.Status)
	}

	return resp, worker, nil
}

func CheckWorkerHealth() {
	WorkerRegistry.Lock()
	defer WorkerRegistry.Unlock()

	for _, worker := range Workers {
		go func(worker *Worker) {
			client := http.Client{Timeout: 1 * time.Second}
			resp, err := client.Get(worker.Address + constants.DispatcherRoutePing)
			worker.LastCheck = time.Now()

			if err != nil || resp.StatusCode != http.StatusOK {
				if worker.Active {
					log.Printf("[HealthCheck] Marking %s as inactive", worker.Address)
				}
				worker.Active = false
			} else {
				if !worker.Active {
					log.Printf("[HealthCheck] Marking %s as active", worker.Address)
				}
				worker.Active = true
			}
			if resp != nil {
				resp.Body.Close()
			}
			worker.LastCheck = time.Now()
		}(worker)
	}
}
