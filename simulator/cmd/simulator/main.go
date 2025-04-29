package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"simulator/internal/constants"
	"simulator/internal/dispatcher"
	"simulator/internal/ipc"
	"simulator/internal/model"
	"simulator/internal/results"
	"simulator/internal/scheduler"
	"simulator/internal/station"
	"simulator/internal/utils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [fcfs|rr]")
		return
	}

	mode := os.Args[1]

	switch mode {
	case "fcfs":
		fmt.Println("Starting in FCFS mode")
	case "rr":
		fmt.Println("Starting in Round Robin mode")
	default:
		fmt.Printf("Invalid mode: %s. Use 'fcfs' or 'rr'.\n", mode)
	}

	quantum := 2 * time.Second

	// Initialize the channels
	ipc.InitChannels(5)

	// Create Mutex for synchronization
	var cutMutex sync.Mutex
	var assembleMutex sync.Mutex
	var packMutex sync.Mutex

	// Initialize the FCFS scheduler
	sched := scheduler.NewFCFS()

	// Creates the stations and defines its order
	go station.Station(constants.StationCutting, ipc.Cutting, ipc.Assembling, &cutMutex, mode, quantum)
	go station.Station(constants.StationAssembling, ipc.Assembling, ipc.Packaging, &assembleMutex, mode, quantum)
	go station.Station(constants.StationPackaging, ipc.Packaging, nil, &packMutex, mode, quantum)

	// Generate 10 products with random arrival times
	go utils.GenerateProducts(sched, 10, 500*time.Millisecond, 2*time.Second)

	// Inserts the products from the scheduler into the first station
	go dispatcher.DispatchProducts(sched, ipc.Cutting)

	var finishedProducts []*model.Product

	for i := 0; i < 10; i++ {
		product := <-ipc.Finished
		finishedProducts = append(finishedProducts, product)
	}

	// All products finished
	results.PrintMetrics(finishedProducts)

	select {}
}
