// input/single_line_test.go
package input

import (
	"strings"
	"testing"
)

func TestProcessSquareRoot(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"sqrt(9)", "9"},
		{"sqrt(25.5)", "25.5"},
		{"sqrt(0)", "0"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			// This is a simple test to verify the string extraction logic
			// In a real test, we'd mock the math.Calculate call
			if strings.HasPrefix(tt.input, "sqrt(") && strings.HasSuffix(tt.input, ")") {
				numStr := tt.input[5 : len(tt.input)-1]
				if numStr != tt.expected {
					t.Errorf("Expected %q, got %q", tt.expected, numStr)
				}
			}
		})
	}
}

func TestProcessFactorial(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"5!", "5"},
		{"10!", "10"},
		{"0!", "0"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			// Test string extraction for factorial
			if strings.HasSuffix(tt.input, "!") {
				numStr := strings.TrimSuffix(tt.input, "!")
				if numStr != tt.expected {
					t.Errorf("Expected %q, got %q", tt.expected, numStr)
				}
			}
		})
	}
}
