package parallelproblems

import (
  "math/rand"
  "runtime"
  "sync"
)

func MonteCarloParallel(n int64) int64 {
    numWorkers := runtime.NumCPU()
    chunk := n / int64(numWorkers)
    results := make(chan int64, numWorkers)
    var wg sync.WaitGroup
    wg.Add(numWorkers)

    for w := 0; w < numWorkers; w++ {
        start := int64(w) * chunk
        end := start + chunk
        if w == numWorkers-1 { end = n }  
        go func(s, e int64) {
            defer wg.Done()
            localRand := rand.New(rand.NewSource(rand.Int63()))
            var inside int64
            for i := s; i < e; i++ {
                x, y := localRand.Float64(), localRand.Float64()
                if x*x+y*y <= 1 { inside++ }
            }
            results <- inside
        }(start, end)
    }

    go func() {
        wg.Wait()
        close(results)
    }()

    var total int64
    for partial := range results {
        total += partial
    }
    return total
}