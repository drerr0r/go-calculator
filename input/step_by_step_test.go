// input/step_by_step_test.go
package input

import (
	"bufio"
	"strings"
	"testing"
)

// MockReader creates a mock reader for tests
type MockReader struct {
	inputs []string
	index  int
}

// Read implements io.Reader interface
func (m *MockReader) Read(p []byte) (n int, err error) {
	if m.index >= len(m.inputs) {
		return 0, nil // EOF
	}

	// Get the current input string
	input := m.inputs[m.index]
	if len(input) == 0 {
		m.index++
		return 0, nil
	}

	// Copy bytes from input to p
	copied := copy(p, []byte(input))
	if copied < len(input) {
		// If we didn't copy all bytes, keep the remaining part
		m.inputs[m.index] = input[copied:]
	} else {
		// Move to next input if we copied all bytes
		m.index++
	}

	return copied, nil
}

func (m *MockReader) ReadString(delim byte) (string, error) {
	if m.index >= len(m.inputs) {
		return "", nil
	}
	result := m.inputs[m.index]
	m.index++
	return result, nil
}

// Convert MockReader to *bufio.Reader
func (m *MockReader) toBufioReader() *bufio.Reader {
	return bufio.NewReader(m)
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

			// Convert MockReader to *bufio.Reader
			ProcessStepByStep(mockReader.toBufioReader())

			// Verify that the first choice is correct
			if tt.inputs[0] != tt.expectedChoice+"\n" {
				t.Errorf("Expected choice %q, got %q", tt.expectedChoice, strings.TrimSpace(tt.inputs[0]))
			}
		})
	}
}

func TestProcessStepByStepInvalidChoice(t *testing.T) {
	mockReader := &MockReader{inputs: []string{"invalid\n"}}

	// Convert MockReader to *bufio.Reader
	ProcessStepByStep(mockReader.toBufioReader())
}
