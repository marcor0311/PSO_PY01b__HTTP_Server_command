package results

import (
	"fmt"
	"time"

	"simulator/internal/model"
)

/**
 * Calculates total time
 * 
 * @param {Product} product - The product to analyze
 * @return {time.Duration} Total time
 */
func TotalTime(product *model.Product) time.Duration {
	return product.ExitedPackage.Sub(product.ArrivalTime)
}

/**
 * Calculates total processing in all stations
 * 
 * @param {Product} product - The product to analyze
 * @return {time.Duration} Total time
 */
func TotalProcessingTime(product *model.Product) time.Duration {
	cutTime := product.ExitedCut.Sub(product.EnteredCut)
	assembleTime := product.ExitedAssemble.Sub(product.EnteredAssemble)
	packageTime := product.ExitedPackage.Sub(product.EnteredPackage)
	return cutTime + assembleTime + packageTime
}

/**
 * Calculates waiting time for a product
 * 
 * Waiting Time = Total Time Time - Processing Time
 * 
 * @param {Product} product - The product to analyze
 * @return {time.Duration} Total waiting time
 */
func WaitingTime(product *model.Product) time.Duration {
	return TotalTime(product) - TotalProcessingTime(product)
}

/**
 * Prints detailed metrics for a list of products.
 *
 * @param {[]*model.Product} products - List of products to print metrics for.
 */
func PrintMetrics(products []*model.Product) {
	var totalTurnaround time.Duration
	var totalWaiting time.Duration

	for _, product := range products {
		ta := TotalTime(product)
		wait := WaitingTime(product)

		totalTurnaround += ta
		totalWaiting += wait

		fmt.Printf("Product %d | Arrival: %v | Total Time: %v | Waiting: %v\n",
			product.Id, product.ArrivalTime.Format("15:04:05"), ta, wait)
	}
}
