package scheduler

import "simulator/internal/model"

type Scheduler interface {
	Add(product *model.Product)
	Next() *model.Product
}
