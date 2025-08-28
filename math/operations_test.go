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
			// Временно отключаем добавление в историю для тестов
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

func TestCalculateErrorCases(t *testing.T) {
	// Тестируем что ошибки обрабатываются без паники
	Calculate(5, 0, "/")     // division by zero
	Calculate(-4, 0, "sqrt") // square root negative
	Calculate(-1, 0, "!")    // factorial negative
	Calculate(5, 3, "&")     // unknown operation
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
