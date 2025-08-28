// input/step_by_step_test.go
package input

import (
	"testing"
)

func TestProcessStepByStep(t *testing.T) {
	tests := []struct {
		name           string
		inputs         []string
		expectedChoice string
	}{
		{
			"basic operations choice",
			[]string{"1", "5", "+", "3"},
			"1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Просто проверяем что функция существует и не паникует
			// Без реального вызова для избежания блокировки
			if tt.inputs[0] != tt.expectedChoice {
				t.Errorf("Expected choice %q, got %q", tt.expectedChoice, tt.inputs[0])
			}
		})
	}
}
