package main

import (
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
	// Initialize the channels 
	ipc.InitChannels(10)

	// Create Mutex for synchronization
	var cutMutex sync.Mutex
	var assembleMutex sync.Mutex
	var packMutex sync.Mutex

	// Initialize the FCFS scheduler
	sched := scheduler.NewFCFS()

	// Creates the stations and defines its order
	go station.Station(constants.StationCutting, ipc.Cutting, ipc.Assembling, utils.RandomDuration(constants.CuttingMinTime, constants.CuttingMaxTime), &cutMutex)
	go station.Station(constants.StationAssembling, ipc.Assembling, ipc.Packaging, utils.RandomDuration(constants.AssemblingMinTime, constants.AssemblingMaxTime), &assembleMutex)
	go station.Station(constants.StationPackaging, ipc.Packaging, nil, utils.RandomDuration(constants.PackagingMinTime, constants.PackagingMaxTime), &packMutex)

	// Generate 10 products with random arrival times
	go utils.GenerateProducts(sched, 10, 500*time.Millisecond, 2*time.Second)

	// Inserts the products from the scheduler into the first station 
	go dispatcher.DispatchProducts(sched, ipc.Cutting)

	var finishedProducts []*model.Product

	for i := 0; i < 10; i++ {
		p := <-ipc.Finished
		finishedProducts = append(finishedProducts, p)
	}

	// All products finished
	results.PrintMetrics(finishedProducts)

	select {}
}
