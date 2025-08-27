// utils/helpers.go
package utils

import "strings"

// AddSpacesAroundOperators добавляет пробелы вокруг операторов
func AddSpacesAroundOperators(input string) string {
	operators := []string{"+", "-", "/", "*", "^", "%", "!", "x", "X", "х", "Х"}

	for _, op := range operators {
		input = strings.ReplaceAll(input, op, " "+op+" ")
	}
	return input
}
