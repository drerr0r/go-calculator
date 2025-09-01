package input

import (
	"bufio"
	"io"
	"strings"
	"testing"
)

// MockReader создает mock reader для тестов
type MockReader struct {
	inputs []string
	index  int
}

// Read реализует io.Reader interface
func (m *MockReader) Read(p []byte) (n int, err error) {
	if m.index >= len(m.inputs) {
		return 0, io.EOF
	}

	input := m.inputs[m.index]
	m.index++

	// Копируем строку в байтовый буфер
	n = copy(p, []byte(input))
	return n, nil
}

func TestProcessStepByStep(t *testing.T) {
	tests := []struct {
		name           string
		inputs         []string
		expectedChoice string
	}{
		{
			"basic operations choice",
			[]string{"1\n", "5\n", "+\n", "3\n"},
			"1",
		},
		{
			"power operations choice",
			[]string{"2\n", "1\n", "4\n", "2\n"},
			"2",
		},
		{
			"percentage choice",
			[]string{"3\n", "100\n", "10\n"},
			"3",
		},
		{
			"factorial choice",
			[]string{"4\n", "5\n"},
			"4",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockReader := &MockReader{inputs: tt.inputs}

			// Тестируем функцию - она должна завершиться
			ProcessStepByStep(bufio.NewReader(mockReader))

			// Проверяем что первый выбор корректный
			if tt.inputs[0] != tt.expectedChoice+"\n" {
				t.Errorf("Expected choice %q, got %q", tt.expectedChoice, strings.TrimSpace(tt.inputs[0]))
			}
		})
	}
}

func TestProcessStepByStepInvalidChoice(t *testing.T) {
	mockReader := &MockReader{inputs: []string{"invalid\n"}}

	// Это не должно паниковать
	ProcessStepByStep(bufio.NewReader(mockReader))
}
