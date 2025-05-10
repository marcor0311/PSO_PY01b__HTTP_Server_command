package handlers

import (
	"errors"
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

// /random?count=n&min=a&max=b: Genera n números aleatorios en el rango indicado.
func Random(count, min, max int) ([]int, error) {
	if min > max {
		return nil, errors.New("min no puede ser mayor que max")
	}
	if count <= 0 {
		return nil, errors.New("count debe ser mayor que cero")
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	numbers := make([]int, count)

	for i := 0; i < count; i++ {
		numbers[i] = r.Intn(max-min+1) + min
	}

	return numbers, nil
}
