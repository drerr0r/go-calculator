// history/history_test.go
package history

import (
	"testing"
)

func TestAddToHistory(t *testing.T) {
	// Clear history before test
	History = []string{}

	tests := []struct {
		expression string
		expected   int
	}{
		{"2 + 2 = 4", 1},
		{"3 * 4 = 12", 2},
		{"10 / 2 = 5", 3},
	}

	for i, tt := range tests {
		t.Run(tt.expression, func(t *testing.T) {
			AddToHistory(tt.expression)

			if len(History) != i+1 {
				t.Errorf("Expected history length %d, got %d", i+1, len(History))
			}

			if History[i] != tt.expression {
				t.Errorf("Expected %q, got %q", tt.expression, History[i])
			}
		})
	}
}

func TestShowHistoryEmpty(t *testing.T) {
	// Clear history
	History = []string{}

	// This test mainly verifies that ShowHistory doesn't panic when empty
	// We can't easily capture stdout in unit tests without refactoring
	ShowHistory()
}
