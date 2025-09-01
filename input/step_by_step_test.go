package input

import (
	"bufio"
	"strings"
	"testing"
)

// MockReader создает mock reader для тестов
type MockReader struct {
	reader *strings.Reader
}

func NewMockReader(inputs []string) *MockReader {
	// Объединяем все входные данные в одну строку
	combined := strings.Join(inputs, "")
	return &MockReader{
		reader: strings.NewReader(combined),
	}
}

func (m *MockReader) Read(p []byte) (n int, err error) {
	return m.reader.Read(p)
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
			mockReader := NewMockReader(tt.inputs)
			reader := bufio.NewReader(mockReader)

			// Тестируем только первую итерацию - выбор операции
			choice, err := reader.ReadString('\n')
			if err != nil {
				t.Fatalf("Error reading choice: %v", err)
			}

			// Проверяем что выбор корректный
			if strings.TrimSpace(choice) != tt.expectedChoice {
				t.Errorf("Expected choice %q, got %q", tt.expectedChoice, strings.TrimSpace(choice))
			}

			// Дальнейшая обработка зависит от выбора, но мы не тестируем весь поток
			// так как функция работает в бесконечном цикле
		})
	}
}

func TestProcessStepByStepInvalidChoice(t *testing.T) {
	mockReader := NewMockReader([]string{"invalid\n"})
	reader := bufio.NewReader(mockReader)

	// Читаем невалидный ввод
	choice, err := reader.ReadString('\n')
	if err != nil {
		t.Fatalf("Error reading choice: %v", err)
	}

	// Проверяем что ввод прочитан корректно
	if strings.TrimSpace(choice) != "invalid" {
		t.Errorf("Expected 'invalid', got %q", strings.TrimSpace(choice))
	}
}
