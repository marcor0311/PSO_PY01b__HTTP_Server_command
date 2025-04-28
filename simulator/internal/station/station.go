package station

import (
	"fmt"
	"time"

	"simulator/internal/constants"
	"simulator/internal/model"
)

/**
 * Processing station
 *
 * @param {string} name - Name of the station
 * @param {<-chan *model.Product} input - Input channel
 * @param {chan<- *model.Product} output - Output channel
 * @param {time.Duration} processingTime - Duration of the station
 */
func Station(name string, input <-chan *model.Product, output chan<- *model.Product, processingTime time.Duration) {
	for product := range input {
		fmt.Printf("Start %s product with id %d\n", name, product.Id)

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
		}

		fmt.Printf("Finished %s product with id %d\n", name, product.Id)

		if output != nil {
			output <- product
		}
	}
}
