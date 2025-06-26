package router

import (
	"bufio"
	"encoding/json"
	"fmt"
	"httpserver/internal/constants"
	"httpserver/internal/worker"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func HandleParallelWordCount(fileURL string) (string, error) {
	const chunkSizeBytes = 15 * 1024 * 1024 // 15MB per chunk

	// Download the file from URL
	resp, err := http.Get(fileURL)
	if err != nil {
		return "", fmt.Errorf("Failed to download file")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Download failed")
	}

	reader := bufio.NewReader(resp.Body)
	globalFrequencies := map[string]int{}
	var mutex sync.Mutex
	var wg sync.WaitGroup

	chunkID := 0

	for {
		chunkBuffer := make([]byte, chunkSizeBytes)
		bytesRead, readErr := io.ReadFull(reader, chunkBuffer)

		if bytesRead > 0 {
			for readErr == nil {
				nextByte, _ := reader.Peek(1)
				if len(nextByte) == 0 || nextByte[0] == ' ' || nextByte[0] == '\n' || nextByte[0] == '\r' {
					break
				}
				extraByte, _ := reader.ReadByte()
				chunkBuffer = append(chunkBuffer, extraByte)
			}

			chunkText := string(chunkBuffer[:bytesRead])
			chunkID++
			wg.Add(1)

			go func(id int, chunk string) {
				defer wg.Done()

				for {
					wordFreq, err := sendChunkToWorker(id, chunk)
					if err != nil {
						log.Printf("[Dispatcher] Failed to process chunk %d: %v", id, err)
						log.Printf("[Dispatcher] Retrying chunk %d with another worker...", id)
						time.Sleep(300 * time.Millisecond)
						continue
					}

					// Merge frequencies into global map
					mutex.Lock()
					for word, count := range wordFreq {
						globalFrequencies[word] += count
					}
					mutex.Unlock()
					break
				}
			}(chunkID, chunkText)
		}

		if readErr == io.EOF {
			break
		}
		if readErr != nil && readErr != io.ErrUnexpectedEOF {
			return "", fmt.Errorf("Error reading the file")
		}
	}

	wg.Wait()

	resultJSON, _ := json.MarshalIndent(globalFrequencies, "", "  ")
	return string(resultJSON), nil
}

func HandleParallelMontecarlo(points int) (string, error) {
	worker.WorkerRegistry.Lock()
	numWorkers := len(worker.Workers)
	worker.WorkerRegistry.Unlock()
	if numWorkers == 0 {
		numWorkers = 1
	}

	pointsPerWorker := points / numWorkers
	remainder := points % numWorkers

	var (
		wg          sync.WaitGroup
		mutex       sync.Mutex
		totalInside int
	)

	for i := 0; i < numWorkers; i++ {
		subPoints := pointsPerWorker
		if i == numWorkers-1 {
			subPoints += remainder
		}
		wg.Add(1)
		go func(id, pts int) {
			defer wg.Done()
			for {
				insideCount, err := sendMontecarloToWorker(id, pts)
				if err != nil {
					log.Printf("[Dispatcher] Monte Carlo chunk %d failed: %v", id, err)
					log.Printf("[Dispatcher] Retrying chunk %d...", id)
					time.Sleep(300 * time.Millisecond)
					continue
				}
				mutex.Lock()
				totalInside += insideCount
				mutex.Unlock()
				break
			}
		}(i, subPoints)
	}

	wg.Wait()

	pi := 4.0 * float64(totalInside) / float64(points)
	result := struct {
		Pi         float64 `json:"pi"`
		Total      int     `json:"total_points"`
		Inside     int     `json:"inside_count"`
		NumWorkers int     `json:"workers"`
	}{
		Pi:         pi,
		Total:      points,
		Inside:     totalInside,
		NumWorkers: numWorkers,
	}

	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	return string(resultJSON), nil
}

func sendMontecarloToWorker(id, points int) (int, error) {
	payload, _ := json.Marshal(struct {
		ID     int `json:"id"`
		Points int `json:"points"`
	}{ID: id, Points: points})

	resp, w, err := worker.SendRequestToWorker(constants.ParallelRouteMontecarlo, string(payload))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		w.MarkInactive()
		return 0, fmt.Errorf("Worker %s returned %s", w.Address, resp.Status)
	}

	var res struct {
		ID     int `json:"id"`
		Inside int `json:"inside"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return 0, err
	}
	log.Printf("[Dispatcher] Received Monte Carlo chunk %d from %s: inside=%d", id, w.Address, res.Inside)
	return res.Inside, nil
}

func sendChunkToWorker(id int, chunk string) (map[string]int, error) {
	payload, _ := json.Marshal(struct {
		ID    int    `json:"id"`
		Chunk string `json:"chunk"`
	}{id, chunk})

	resp, worker, err := worker.SendRequestToWorker(constants.ParallelRouteCount, string(payload))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res struct {
		ID   int            `json:"id"`
		Freq map[string]int `json:"freq"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	worker.Completed++

	log.Printf("[Dispatcher] Successfully received chunk %d from worker %s", id, worker.Address)
	return res.Freq, nil
}
