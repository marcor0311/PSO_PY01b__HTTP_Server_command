package dispatcher

import (
	"time"

	"simulator/internal/model"
	"simulator/internal/scheduler"
)

/**
 * Dispatches products from the scheduler into the first station's input channel.
 *
 * @param {scheduler.Scheduler} scheduler - The scheduler to get the products
 * @param {chan<- *model.Product} inputChannel - Input channel of 1st station
 */
func DispatchProducts(scheduler scheduler.Scheduler, inputChannel chan<- *model.Product) {
	for {
		product := scheduler.Next()
		if product == nil {
			time.Sleep(100 * time.Millisecond)
			continue
		}

		inputChannel <- product
	}
}
