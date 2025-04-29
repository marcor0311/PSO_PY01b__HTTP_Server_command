package scheduler

import (
    "sync"
    "simulator/internal/model"
    "time"
)

/**
 * RoundRobin implements the Scheduler interface using a cyclic queue with quantum-based time slicing.
 */
type RoundRobin struct {
    mutex     sync.Mutex
    queue     []*model.Product
    index     int
    quantum   time.Duration
    remaining map[*model.Product]time.Duration
}

/**
 * NewRoundRobin returns a new RoundRobin scheduler configured with a given quantum and service time.
 *
 * @param {time.Duration} quantum - Duration of each time slice.
 * @param {time.Duration} serviceTime - Total required processing time.
 * @returns {RoundRobin}
 */
func NewRoundRobin(quantum, serviceTime time.Duration) *RoundRobin {
    return &RoundRobin{
        queue:     make([]*model.Product, 0),
        quantum:   quantum,
        remaining: make(map[*model.Product]time.Duration),
    }
}

/**
 * Add places a product in the queue and initializes its remaining time.
 *
 * @param {model.Product} p - The product to be scheduled.
 */
func (rr *RoundRobin) Add(p *model.Product) {
    rr.mutex.Lock()
    defer rr.mutex.Unlock()
    rr.queue = append(rr.queue, p)
    rr.remaining[p] = rr.quantum
}

/**
 * Next returns the next product to be processed, applying quantum-based slicing and rotation.
 *
 * @returns {model.Product | nil}
 */
func (rr *RoundRobin) Next() *model.Product {
    rr.mutex.Lock()
    defer rr.mutex.Unlock()

    if len(rr.queue) == 0 {
        return nil
    }

    product := rr.queue[rr.index]
    rr.index = (rr.index + 1) % len(rr.queue)

    // Simulate quantum-based slicing
    if rr.remaining[product] > 0 {
        rr.remaining[product] -= rr.quantum
    }

    // If finished remove from queue
    if rr.remaining[product] <= 0 {
        rr.queue = append(rr.queue[:rr.index], rr.queue[rr.index+1:]...)
        delete(rr.remaining, product)
        if rr.index > 0 {
            rr.index--
        }
    }

    return product
}
