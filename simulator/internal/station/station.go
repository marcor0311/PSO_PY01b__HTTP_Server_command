package station

import (
	"fmt"
	"sync"
	"time"

	"simulator/internal/constants"
	"simulator/internal/ipc"
	"simulator/internal/model"
)

/**
 * Station simulates a processing station (Cutting, Assembling, Packaging) with synchronization.
 *
 * @param {string} name - The name of the station (constants).
 * @param {<-chan *model.Product} input - Receive-only channel.
 * @param {chan<- *model.Product} output - Send-only channel (nil if last station).
 * @param {time.Duration} processingTime - Time to simulate processing.
 * @param {*sync.Mutex} mutex - Mutex to synchronize access to the station.
 */
 func Station(name string, input <-chan *model.Product, output chan<- *model.Product, processingTime time.Duration, mutex *sync.Mutex) {
	for product := range input {
		mutex.Lock()

		fmt.Printf("[%s] Received product %d\n", name, product.Id)

		now := time.Now()
		switch name {
		case constants.StationCutting:
			product.EnteredCut = now
		case constants.StationAssembling:
			product.EnteredAssemble = now
		case constants.StationPackaging:
			product.EnteredPackage = now
		}

		time.Sleep(processingTime)

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