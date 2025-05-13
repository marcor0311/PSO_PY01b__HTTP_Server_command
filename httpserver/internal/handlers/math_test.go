package handlers_test

import (
	"httpserver/internal/handlers"
	"testing"
)

// Test cases for Fibonacci function
func TestFibonacci_Zero(t *testing.T) {
	_, err := handlers.Fibonacci(0)
	if err == nil {
		t.Error("Expected error for n = 0, got nil")
	}
}

func TestFibonacci_Negative(t *testing.T) {
	_, err := handlers.Fibonacci(-5)
	if err == nil {
		t.Error("Expected error for negative n, got nil")
	}
}

func TestFibonacci_One(t *testing.T) {
	result, err := handlers.Fibonacci(1)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

func TestFibonacci_Recursive(t *testing.T) {
	result, err := handlers.Fibonacci(20)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	expected := 4181 // 20th Fibonacci number
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// Test cases for Random function
func TestRandom_MinGreaterThanMax(t *testing.T) {
	_, err := handlers.Random(3, 10, 5)
	if err == nil {
		t.Error("Expected error for min > max, got nil")
	}
}

func TestRandom_CountZero(t *testing.T) {
	_, err := handlers.Random(0, 1, 10)
	if err == nil {
		t.Error("Expected error for count = 0, got nil")
	}
}

func TestRandom_CountNegative(t *testing.T) {
	_, err := handlers.Random(-1, 1, 10)
	if err == nil {
		t.Error("Expected error for negative count, got nil")
	}
}

func TestRandom_InvalidRangeZeroWidth(t *testing.T) {
	_, err := handlers.Random(5, 5, 4)
	if err == nil {
		t.Error("Expected error for invalid range, got nil")
	}
}

func TestRandom_SingleValueRange(t *testing.T) {
	numbers, err := handlers.Random(3, 7, 7)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	for _, n := range numbers {
		if n != 7 {
			t.Errorf("Expected all numbers to be 7, got %d", n)
		}
	}
}
