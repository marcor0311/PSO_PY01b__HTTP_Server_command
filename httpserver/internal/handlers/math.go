package handlers

import (
	"fmt"
	"math/rand"
	"time"
)

// /fibonacci?num=N: Cálculo recursivo del número N de la serie de Fibonacci.
func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// /random?count=n&min=a&max=b: Returns a list of n random integers between min and max.
func Random(count, min, max int) ([]int, error) {
	if min > max {
		return nil, fmt.Errorf("Invalid range: min (%d) cannot be greater than max (%d)", min, max)
	}
	if count <= 0 {
		return nil, fmt.Errorf("Invalid count: must be greater than zero")
	}
	if (max - min + 1) <= 0 {
		return nil, fmt.Errorf("Ivalid range")
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	numbers := make([]int, count)

	for i := 0; i < count; i++ {
		numbers[i] = r.Intn(max-min+1) + min
	}

	return numbers, nil
}