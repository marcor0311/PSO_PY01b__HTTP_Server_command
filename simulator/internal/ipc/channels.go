package ipc

import "simulator/internal/model"

var (
    Cutting    chan *model.Product
    Assembling chan *model.Product
    Packaging  chan *model.Product
    Finished chan *model.Product
)

func InitChannels(bufferSize int) {
    Cutting = make(chan *model.Product, bufferSize)
    Assembling = make(chan *model.Product, bufferSize)
    Packaging = make(chan *model.Product, bufferSize)
    Finished = make(chan *model.Product, bufferSize)
}
