package main

import (
    "fmt"
    "time"

	"simulator/internal/constants"
	"simulator/internal/model"
	"simulator/internal/station"
)

func main() {

	// inicia canales
	input := make(chan *model.Product, 1)
	output := make(chan *model.Product, 1)

	// producto 
	product := &model.Product{
		Id: 1,
		ArrivalTime: time.Now(),
	}

	// inicia canal y lo envia al canal inicial
	input <- product
	close(input) // cierra el canal

	// ejecuta la estación
	go station.Station(constants.StationCutting, input, output, 2*time.Second)

	// resultado del canal
	result := <-output

	fmt.Printf("Product %d processed!\n", result.Id)
	fmt.Printf("Entered Cut: %v\n", result.EnteredCut)
	fmt.Printf("Exited Cut: %v\n", result.ExitedCut)
}

