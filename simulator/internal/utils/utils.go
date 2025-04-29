package utils

import (
	"math/rand"
	"time"

	"simulator/internal/model"
	"simulator/internal/scheduler"
)

/**
 * Returns a random time between a minimum and maximum
 *
 * @param {time.Duration} min - Minimum duration
 * @param {time.Duration} max - Maximum duration
 * @return {time.Duration} Random duration between min and max
 */
func RandomDuration(min, max time.Duration) time.Duration {
	diff := max - min
	return min + time.Duration(rand.Int63n(int64(diff)))
}

/**
 * Generates products with random arrival times and adds them to the scheduler queue.
 *
 * @param {scheduler.Scheduler} scheduler - Scheduler to add products to.
 * @param {int} totalProducts - Number of products to generate
 * @param {time.Duration} minInterval - Minimun time between arrivals
 * @param {time.Duration} maxInterval - Maximum time between arrivals
 */
func GenerateProducts(scheduler scheduler.Scheduler, totalProducts int, minInterval, maxInterval time.Duration) {
	for i := 1; i <= totalProducts; i++ {
		arrivalDelay := minInterval + time.Duration(rand.Int63n(int64(maxInterval-minInterval)))

		time.Sleep(arrivalDelay) // Simulate arrival time

		product := &model.Product{
			Id:          i,
			ArrivalTime: time.Now(),
		}

		scheduler.Add(product)
	}
}
