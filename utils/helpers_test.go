// utils/helpers_test.go
package utils

import (
	"testing"
)

func TestAddSpacesAroundOperators(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"addition without spaces", "5+3", "5 + 3"},
		{"addition with spaces", "5 + 3", "5 + 3"},
		{"subtraction", "10-4", "10 - 4"},
		{"multiplication", "6*7", "6 * 7"},
		{"division", "15/3", "15 / 3"},
		{"power", "2^3", "2 ^ 3"},
		{"percentage", "100%10", "100 % 10"},
		{"factorial", "5!", "5 !"},
		{"multiple operators", "5+3*2", "5 + 3 * 2"},
		{"with spaces", "5 + 3 * 2", "5 + 3 * 2"},
		{"russian x", "5х3", "5 х 3"},
		{"uppercase X", "5X3", "5 X 3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AddSpacesAroundOperators(tt.input)
			if result != tt.expected {
				t.Errorf("AddSpacesAroundOperators(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
