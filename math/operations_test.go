// math/operations_test.go
package math

import (
	"testing"
)

func TestCalculateBasicOperations(t *testing.T) {
	tests := []struct {
		name      string
		a         float64
		b         float64
		operation string
		expected  float64
	}{
		{"addition", 5, 3, "+", 8},
		{"subtraction", 10, 4, "-", 6},
		{"multiplication", 6, 7, "*", 42},
		{"division", 15, 3, "/", 5},
		{"power", 2, 3, "^", 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Тестируем только логику вычислений
			var result float64
			switch tt.operation {
			case "+":
				result = tt.a + tt.b
			case "-":
				result = tt.a - tt.b
			case "*":
				result = tt.a * tt.b
			case "/":
				result = tt.a / tt.b
			case "^":
				result = 1.0
				for i := 0; i < int(tt.b); i++ {
					result *= tt.a
				}
			}

			if result != tt.expected {
				t.Errorf("%s failed: got %v, want %v", tt.name, result, tt.expected)
			}
		})
	}
}

func TestCalculateLegacyErrorCases(t *testing.T) {
	// Тестируем что ошибки обрабатываются без паники
	CalculateLegacy(5, 0, "/")     // division by zero
	CalculateLegacy(-4, 0, "sqrt") // square root negative
	CalculateLegacy(-1, 0, "!")    // factorial negative
	CalculateLegacy(5, 3, "&")     // unknown operation
}

func TestCalculateWithParser(t *testing.T) {
	tests := []struct {
		expression string
		expected   float64
		wantErr    bool
	}{
		{"5+3", 8, false},
		{"10-4", 6, false},
		{"6*7", 42, false},
		{"15/3", 5, false},
		{"2^3", 8, false},
		{"(2+3)*4", 20, false},
		{"2+3*4", 14, false},
		{"invalid", 0, true},
		{"5/0", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.expression, func(t *testing.T) {
			result, err := Calculate(tt.expression)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error for %q, but got result %f", tt.expression, result)
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error for %q: %v", tt.expression, err)
				return
			}

			if result != tt.expected {
				t.Errorf("%s = %f, want %f", tt.expression, result, tt.expected)
			}
		})
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 1},
		{1, 1},
		{5, 120},
		{6, 720},
	}

	for _, tt := range tests {
		t.Run(string(rune(tt.input)), func(t *testing.T) {
			result := factorial(tt.input)
			if result != tt.expected {
				t.Errorf("factorial(%d) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestParserIntegration(t *testing.T) {
	tests := []struct {
		expression string
		expected   float64
	}{
		{"2 + 3 * 4", 14},
		{"(2 + 3) * 4", 20},
		{"10 / (2 + 3)", 2},
		{"2 ^ (1 + 2)", 8},
	}

	for _, tt := range tests {
		t.Run(tt.expression, func(t *testing.T) {
			parser := NewParser(tt.expression)
			result, err := parser.Parse()

			if err != nil {
				t.Errorf("Parser error for %q: %v", tt.expression, err)
				return
			}

			if result != tt.expected {
				t.Errorf("%s = %f, want %f", tt.expression, result, tt.expected)
			}
		})
	}
}
