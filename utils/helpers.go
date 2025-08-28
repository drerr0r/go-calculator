// utils/helpers.go
package utils

import (
	"regexp"
	"strings"
)

// AddSpacesAroundOperators добавляет пробелы вокруг операторов для лучшего отображения
func AddSpacesAroundOperators(expression string) string {
	// Создаем карту операторов для обработки
	operatorsMap := map[string]bool{
		"+": true, "-": true, "*": true, "/": true,
		"!": true, "^": true, "%": true, "х": true, "X": true,
	}

	// Проходим по строке и добавляем пробелы вокруг операторов
	var result strings.Builder
	for i, char := range expression {
		currentChar := string(char)

		// Если текущий символ - оператор
		if operatorsMap[currentChar] {
			// Добавляем пробел перед оператором, если предыдущий символ не пробел и не начало строки
			if i > 0 && expression[i-1] != ' ' {
				result.WriteString(" ")
			}
			result.WriteString(currentChar)
			// Добавляем пробел после оператора, если следующий символ не пробел и не конец строки
			if i < len(expression)-1 && expression[i+1] != ' ' {
				result.WriteString(" ")
			}
		} else {
			result.WriteString(currentChar)
		}
	}

	// Убираем лишние пробелы
	finalResult := result.String()
	finalResult = strings.TrimSpace(finalResult)
	spaceRe := regexp.MustCompile(`\s+`)
	finalResult = spaceRe.ReplaceAllString(finalResult, " ")

	return finalResult
}
