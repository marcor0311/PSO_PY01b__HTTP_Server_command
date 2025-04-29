package scheduler

import (
	"sync"

	"simulator/internal/model"
)

type FCFS struct {
	queue []*model.Product
	lock  sync.Mutex
}

/**
 * Creates and returns a new FCFS scheduler.
 *
 * @return {Scheduler} - New instance of FCFS Scheduler.
 */
func NewFCFS() *FCFS {
	return &FCFS{
		queue: make([]*model.Product, 0),
	}
}

/**
 * Adds a product to the scheduler queue.
 *
 * @param {Product} product - The product to schedule.
 */
func (s *FCFS) Add(product *model.Product) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.queue = append(s.queue, product)
}

/**
 * Returns the next product to dispatch, or null if none are ready.
 *
 * @return {Product} The next scheduled product, or null.
 */
func (s *FCFS) Next() *model.Product {
	s.lock.Lock()
	defer s.lock.Unlock()

	if len(s.queue) == 0 {
		return nil
	}
	next := s.queue[0]
	s.queue = s.queue[1:]
	return next
}
