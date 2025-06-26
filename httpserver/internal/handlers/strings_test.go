package handlers_test

import (
	"reflect"
	"strings"
	"testing"

	"httpserver/internal/handlers"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"abc", "cba"},
		{"", ""},
		{"a", "a"},
		{"hello world", "dlrow olleh"},
	}

	for _, test := range tests {
		result := handlers.ReverseString(test.input)
		if result != test.expected {
			t.Errorf("ReverseString(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestToUpper(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"abc", "ABC"},
		{"AbC123", "ABC123"},
	}

	for _, test := range tests {
		result := handlers.ToUpper(test.input)
		if result != test.expected {
			t.Errorf("ToUpper(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestHashSHA256(t *testing.T) {
	tests := []struct {
		input    string
		expected string 
	}{
		{"abc", ""},
		{"123", ""},
		{"Hello, world!", ""},
	}

	for _, test := range tests {
		result := handlers.HashSHA256(test.input)
		if len(result) != 64 {
			t.Errorf("HashSHA256(%q) returned string of length %d, expected 64", test.input, len(result))
		}
		if !isHex(result) {
			t.Errorf("HashSHA256(%q) returned non-hex string: %q", test.input, result)
		}
	}
}

func isHex(s string) bool {
	for _, r := range s {
		if !strings.ContainsRune("0123456789abcdef", r) {
			return false
		}
	}
	return true
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{
			name:  "simple sentence",
			input: "Hello world hello",
			expected: map[string]int{
				"hello": 2,
				"world": 1,
			},
		},
		{
			name:  "punctuation and case",
			input: "Go, go, GO!",
			expected: map[string]int{
				"go": 3,
			},
		},
		{
			name:  "mixed words",
			input: "One fish two fish red fish blue fish.",
			expected: map[string]int{
				"one":  1,
				"fish": 4,
				"two":  1,
				"red":  1,
				"blue": 1,
			},
		},
		{
			name:     "empty input",
			input:    "",
			expected: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := handlers.CountWords(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %+v, got %+v", tt.expected, result)
			}
		})
	}
}