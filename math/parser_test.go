// math/parser_test.go
package math

import (
	"testing"
)

func TestParser(t *testing.T) {
	tests := []struct {
		expression string
		expected   float64
		wantErr    bool
	}{
		// Базовые операции
		{"2+3", 5, false},
		{"5-2", 3, false},
		{"4*3", 12, false},
		{"10/2", 5, false},

		// Скобки
		{"(2+3)*4", 20, false},
		{"2*(3+4)", 14, false},
		{"(1+2)*(3+4)", 21, false},
		{"10/(2+3)", 2, false},

		// Вложенные скобки
		{"((2+3)*4)", 20, false},
		{"(2*(3+4))", 14, false},

		// Приоритет операций
		{"2+3*4", 14, false},
		{"(2+3)*4", 20, false},
		{"2*3+4", 10, false},

		// Степени
		{"2^3", 8, false},
		{"2^3+1", 9, false},
		{"2^(3+1)", 16, false},

		// Ошибки
		{"2+", 0, true},
		{"(2+3", 0, true},
		{"2+*3", 0, true},
		{"abc", 0, true},
		{")2+3(", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.expression, func(t *testing.T) {
			parser := NewParser(tt.expression)
			result, err := parser.Parse()

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

func TestPowFunction(t *testing.T) {
	tests := []struct {
		base     float64
		exponent float64
		expected float64
	}{
		{2, 3, 8},
		{3, 2, 9},
		{5, 0, 1},
		{2, 1, 2},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := Pow(tt.base, tt.exponent)
			if result != tt.expected {
				t.Errorf("Pow(%.1f, %.1f) = %.1f, want %.1f", tt.base, tt.exponent, result, tt.expected)
			}
		})
	}
}
