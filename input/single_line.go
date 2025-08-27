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
	fmt.Print("Введите выражение (например: 5 + 3, 4^2, sqrt(9)): ")
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

	if strings.HasSuffix(input, "!") {
		processFactorial(input)
		return
	}

	// Обработка обычных операций
	processedInput := utils.AddSpacesAroundOperators(input)
	parts := strings.Fields(processedInput)

	// Для операций с одним числом (например 5!)
	if len(parts) == 2 && parts[1] == "!" {
		processFactorial(strings.Join(parts, " "))
		return
	}

	if len(parts) != 3 {
		fmt.Println("Ошибка: Введите выражение в формате 'число оператор число'")
		fmt.Println("Примеры: 5 + 3, 4^2, 10%")
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

	math.Calculate(a, b, operation)
}

func processSquareRoot(input string) {
	numStr := input[5 : len(input)-1]
	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		fmt.Println("Ошибка: неверное число внутри sqrt()")
		return
	}
	math.Calculate(num, 0, "sqrt")
}

func processFactorial(input string) {
	numStr := strings.TrimSuffix(input, "!")
	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		fmt.Println("Ошибка: неверное число для факториала")
		return
	}
	math.Calculate(num, 0, "!")
}
