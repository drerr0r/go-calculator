// input/single_line.go
package input

import (
	"bufio"
	"calculator/math"
	"calculator/utils"
	"fmt"
	"strconv"
	"strings"
)

// ProcessSingleLine обрабатывает ввод одной строкой
func ProcessSingleLine(reader *bufio.Reader) {
	fmt.Print("Введите выражение (например: 5 + 3, (2+3)*4, sqrt(9)): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения ввода:", err)
		return
	}
	input = strings.TrimSpace(input)

	// Обработка специальных операций
	if strings.HasPrefix(input, "sqrt(") && strings.HasSuffix(input, ")") {
		processSquareRoot(input)
		return
	}

	if strings.HasSuffix(input, "!") && !containsOperatorsExceptFactorial(input) {
		processFactorial(input)
		return
	}

	// Проверяем, содержит ли выражение скобки или сложные операции
	if containsParentheses(input) || isComplexExpression(input) {
		processComplexExpression(input)
		return
	}

	// Обработка простых операций без скобок
	processedInput := utils.AddSpacesAroundOperators(input)
	parts := strings.Fields(processedInput)

	// Для операций с одним числом (например 5!)
	if len(parts) == 2 && parts[1] == "!" {
		processFactorial(strings.Join(parts, " "))
		return
	}

	if len(parts) != 3 {
		fmt.Println("Ошибка: Введите выражение в формате 'число оператор число'")
		fmt.Println("Примеры: 5 + 3, 4^2, 10%, (2+3)*4")
		return
	}

	a, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		fmt.Println("Ошибка: первое значение не является числом")
		return
	}

	operation := parts[1]

	b, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		fmt.Println("Ошибка: второе значение не является числом")
		return
	}

	math.CalculateLegacy(a, b, operation)
}

// containsParentheses проверяет, содержит ли строка скобки
func containsParentheses(input string) bool {
	return strings.Contains(input, "(") || strings.Contains(input, ")")
}

// isComplexExpression проверяет, является ли выражение сложным
func isComplexExpression(input string) bool {
	// Считаем количество операторов
	operators := []string{"+", "-", "*", "/", "^", "%"}
	operatorCount := 0

	for _, op := range operators {
		operatorCount += strings.Count(input, op)
	}

	// Если больше одного оператора - сложное выражение
	return operatorCount > 1
}

// containsOperatorsExceptFactorial проверяет наличие операторов кроме факториала
func containsOperatorsExceptFactorial(input string) bool {
	operators := []string{"+", "-", "*", "/", "^", "%", "("}
	for _, op := range operators {
		if strings.Contains(input, op) {
			return true
		}
	}
	return false
}

// processComplexExpression обрабатывает сложные выражения со скобками
func processComplexExpression(input string) {
	result, err := math.Calculate(input)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		fmt.Println("Примеры корректных выражений:")
		fmt.Println("  (5 + 3) * 2")
		fmt.Println("  2 ^ (3 + 1)")
		fmt.Println("  10 / (2 + 3)")
		fmt.Println("  2 + 3 * 4")
		return
	}

	fmt.Printf("Результат: %s = %.2f\n",
		utils.AddSpacesAroundOperators(input), result)
}

func processSquareRoot(input string) {
	numStr := input[5 : len(input)-1]
	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		fmt.Println("Ошибка: неверное число внутри sqrt()")
		return
	}
	math.CalculateLegacy(num, 0, "sqrt")
}

func processFactorial(input string) {
	numStr := strings.TrimSuffix(input, "!")
	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		fmt.Println("Ошибка: неверное число для факториала")
		return
	}
	math.CalculateLegacy(num, 0, "!")
}
