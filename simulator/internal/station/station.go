package station

import (
	"fmt"
	"sync"
	"time"

	"simulator/internal/constants"
	"simulator/internal/ipc"
	"simulator/internal/model"
	"simulator/internal/utils"
)

/**
 * Station simulates a processing station (Cutting, Assembling, Packaging) with synchronization.
 *
 * @param {string} name - The name of the station (constants).
 * @param {<-chan *model.Product} input - Receive-only channel.
 * @param {chan<- *model.Product} output - Send-only channel (nil if last station).
 * @param {*sync.Mutex} mutex - Mutex to synchronize access to the station.
 * @param {string} algorithm - Scheduling algorithm ("fcfs" or "rr").
 * @param {time.Duration} quantum - Time slice for round-robin.
 */
func Station(name string, input <-chan *model.Product, output chan<- *model.Product, mutex *sync.Mutex, algorithm string, quantum time.Duration) {
	queue := []*model.Product{}

	for {
		var product *model.Product

		var processingTime time.Duration

		switch name {
		case constants.StationCutting:
			processingTime = utils.RandomDuration(constants.CuttingMinTime, constants.CuttingMaxTime)
		case constants.StationAssembling:
			processingTime = utils.RandomDuration(constants.AssemblingMinTime, constants.AssemblingMaxTime)
		case constants.StationPackaging:
			processingTime = utils.RandomDuration(constants.PackagingMinTime, constants.PackagingMaxTime)
		}

		select {
		case incoming, ok := <-input:
			if !ok {
				return
			}
			product = incoming
		default:
			if len(queue) > 0 {
				product = queue[0]
				queue = queue[1:]
			} else {
				time.Sleep(10 * time.Millisecond)
				continue
			}
		}

		mutex.Lock()

		fmt.Printf("[%s] Processing product %d\n", name, product.Id)

		now := time.Now()
		switch name {
		case constants.StationCutting:
			product.EnteredCut = now
		case constants.StationAssembling:
			product.EnteredAssemble = now
		case constants.StationPackaging:
			product.EnteredPackage = now
		}

		if algorithm == "rr" {
			// Initialize RemainingTime if it's a new product
			if product.RemainingTime == 0 {
				product.RemainingTime = processingTime
			}

			slice := quantum
			if product.RemainingTime < quantum {
				slice = product.RemainingTime
			}
			fmt.Printf("[%s] Processing product %d for %v\n", name, product.Id, slice)
			time.Sleep(slice)
			product.RemainingTime -= slice

			// Requeue unfinished products internally instead of sending them back to output channel
			if product.RemainingTime > 0 {
				fmt.Printf("[%s] Re-enqueued product %d with %v remaining\n", name, product.Id, product.RemainingTime)
				queue = append(queue, product)
				mutex.Unlock()
				continue
			}
		} else {
			fmt.Printf("[%s] Processing product %d for %v\n", name, product.Id, processingTime)
			time.Sleep(processingTime)
		}

		now = time.Now()
		switch name {
		case constants.StationCutting:
			product.ExitedCut = now
		case constants.StationAssembling:
			product.ExitedAssemble = now
		case constants.StationPackaging:
			product.ExitedPackage = now
			ipc.Finished <- product
		}

		fmt.Printf("[%s] Finished product %d\n", name, product.Id)

		mutex.Unlock()

		if output != nil {
			output <- product
		}
	}
}