package handlers_test

import (
	"httpserver/internal/handlers"
	"testing"
)

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
