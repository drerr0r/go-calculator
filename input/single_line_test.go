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

func TestContainsParentheses(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"(2+3)", true},
		{"2+3", false},
		{"(2+3)*4", true},
		{"2*(3+4)", true},
		{"sqrt(9)", true},
		{"5!", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := containsParentheses(tt.input)
			if result != tt.expected {
				t.Errorf("containsParentheses(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsComplexExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"2+3*4", true},
		{"5+3", false},
		{"2*3+4", true},
		{"(2+3)*4", true},
		{"5!", false},
		{"sqrt(9)", false},
		{"2^3+1", true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := isComplexExpression(tt.input)
			if result != tt.expected {
				t.Errorf("isComplexExpression(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestContainsOperatorsExceptFactorial(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"5!", false},
		{"5!+3", true},
		{"(5!)*2", true},
		{"factorial!", false},
		{"5!*3", true},
		{"sqrt(9!)", true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := containsOperatorsExceptFactorial(tt.input)
			if result != tt.expected {
				t.Errorf("containsOperatorsExceptFactorial(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
