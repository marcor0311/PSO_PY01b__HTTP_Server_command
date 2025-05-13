package handlers

import (
	"fmt"
	"math/rand"
	"time"
)

// /fibonacci?num=N: Cálculo recursivo del número N de la serie de Fibonacci.
func Fibonacci(n int) (int, error) {
    if n <= 0 {
        return 0, fmt.Errorf("Invalid input: n (%d) must be greater than zero", n)
    }

    if n == 1 {
        return 0, nil
    }
    if n == 2 {
        return 1, nil
    }

    prev, err1 := Fibonacci(n - 1)
    if err1 != nil {
        return 0, err1
    }
    prevPrev, err2 := Fibonacci(n - 2)
    if err2 != nil {
        return 0, err2
    }

    return prev + prevPrev, nil
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