package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var history []string

func main() {
	fmt.Println("Добро пожаловать в продвинутый калькулятор на Go!")

	// Создаем читателя для ввода консоли
	reader := bufio.NewReader(os.Stdin)

	for {
		// Запрос на ввод
		fmt.Print("\nВыберите режим ввода:\n1 - Одна строка (например: 5 + 3)\n2 - Пошаговый ввод\nистория - Показать историю вычислений\nстоп - Выход\nВаш выбор: ")
		choice, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения ввода:", err)
			continue
		}

		choice = strings.TrimSpace(choice)

		if strings.ToLower(choice) == "стоп" {
			fmt.Print("До свидания!")
			break
		}

		if strings.ToLower(choice) == "история" {
			showHistory()
			continue
		}

		switch choice {
		case "1":
			processSingleLineInput(reader)
		case "2":
			processStepByStepInput(reader)
		default:
			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
	}
}

func processSingleLineInput(reader *bufio.Reader) {
	fmt.Print("Введите выражение (например 5 + 3, 4^2, sqrt(9))): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения ввода:", err)
		return
	}
	input = strings.TrimSpace(input)

	//Обррабатываем специальные оппеации: sqrt, !, ^
	if strings.HasPrefix(input, "sqrt(") && strings.HasPrefix(input, ")") {
		processSquareRoot(input)
		return
	}

	if strings.HasSuffix(input, "!") {
		processFactorial(input)
	}

	// Обрабатываем обычные операции
	processedInput := addSpacesAroundOperators(input)

	// Разбиваем строку на части
	parts := strings.Fields(processedInput)
	if len(parts) != 3 {
		fmt.Println("Ошибка: Введите выражение в формате 'число оператор число'")
		fmt.Println("Пример: 5 + 3, 4^2, 10%, sqrt(9), 5!")
		return
	}

	//Для операций с одним числом (например 5!)
	if len(parts) == 2 && parts[1] == "!" {
		processFactorial(strings.Join(parts, " "))
		return
	}

	if len(parts) != 3 {
		fmt.Println("Ошибка: Введите выражение в формате 'число оператор число'")
		fmt.Println("Примеры: 5 + 3, 4^2, 10%")
	}

	// Парсим первое число
	a, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		fmt.Println("Ошибка: первое значение не является числом")
		return
	}

	// Получаем оператор
	operation := parts[1]

	// Парсим второе число
	b, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		fmt.Println("Ошибка: Второе значение не является числом")
		return
	}

	// Вычисляем результат
	calculateAndPrint(a, b, operation)
}

func addSpacesAroundOperators(input string) string {
	operators := []string{"+", "-", "/", "*", "x", "X", "х", "Х", "^", "%", "!"}

	for _, op := range operators {
		input = strings.ReplaceAll(input, op, " "+op+" ")
	}
	return input
}

func processStepByStepInput(reader *bufio.Reader) {
	fmt.Print("Выбереите операцию:/n1 - Базовые (+, -, *, /)\n2 - Степень и корень (^, sqrt)\n3 - Проценты (%)\n4 - Факториал (!)\nВаш выбор: ")

	opType, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения ввода:", err)
		return
	}
	opType = strings.TrimSpace(opType)

	switch opType {
	case "1":
		processBasicOperations(reader)
	case "2":
		processPowerOperations(reader)
	case "3":
		processPercentage(reader)
	case "4":
		processFactorialInput(reader)
	default:
		fmt.Println("Неверный выбор.")
	}

}

func processBasicOperations(reader *bufio.Reader) {
	var a, b float64
	var operation string

	fmt.Print("Введите первое число: ")
	inputA, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения ввода:", err)
		return
	}
	inputA = strings.TrimSpace(inputA)
	_, err = fmt.Sscanf(inputA, "%f", &a)
	if err != nil {
		fmt.Println("Ошибка чтения числа:", err)
		return
	}

	fmt.Print("Введите оперрацию (+, -, *, /): ")
	inputOp, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения ввода:", err)
		return
	}
	operation = strings.TrimSpace(inputOp)

	fmt.Print("Введите второе число: ")
	inputB, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения ввода:", err)
		return
	}
	inputB = strings.TrimSpace(inputB)
	_, err = fmt.Sscanf(inputB, "%f", &a)
	if err != nil {
		fmt.Println("Ошибка чтения числа:", err)
		return
	}

	calculateAndPrint(a, b, operation)
}

func processPowerOperations(reader *bufio.Reader) {
	fmt.Print("Выберите:\n1 - Возведение в степень (a^b)\n2 - Квадратный корень (sqrt(a))\nВаш выбор: ")

	choice, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения ввода:", err)
		return
	}
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		var a, b float64
		fmt.Print("Введите число: ")
		inputA, _ := reader.ReadString('\n')
		inputA = strings.TrimSpace(inputA)
		fmt.Println(inputA, "%f", &a)

		fmt.Print("Введите степень: ")
		inputB, _ := reader.ReadString('\n')
		inputB = strings.TrimSpace(inputB)
		fmt.Println(inputB, "%f", &b)

		calculateAndPrint(a, b, "^")

	case "2":
		var a float64
		fmt.Print("Введите число: ")
		inputA, _ := reader.ReadString('\n')
		inputA = strings.TrimSpace(inputA)
		fmt.Println(inputA, "%f", &a)

		calculateAndPrint(a, 0, "sqrt")
	}
}

func processPercentage(reader *bufio.Reader) {

	var a, b float64
	fmt.Print("Введите число: ")
	inputA, _ := reader.ReadString('\n')
	inputA = strings.TrimSpace(inputA)
	fmt.Println(inputA, "%f", &a)

	fmt.Print("Введите процент: ")
	inputB, _ := reader.ReadString('\n')
	inputB = strings.TrimSpace(inputB)
	fmt.Println(inputB, "%f", &b)

	calculateAndPrint(a, b, "%")
}

func processFactorialInput(reader *bufio.Reader) {
	var a float64

	fmt.Print("Введите число для вычисления фактоиала: ")
	inputA, _ := reader.ReadString('\n')
	inputA = strings.TrimSpace(inputA)
	fmt.Println(inputA, "%f", a)

	calculateAndPrint(a, 0, "!")
}

func processSquareRoot(input string) {
	//Извлекаем число из sqrt(9) -> 9
	numStr := input[5 : len(input)-1]
	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		fmt.Println("Ошибка: невеное число внутри sqrt()")
		return
	}
	calculateAndPrint(num, 0, "sqrt")

}

func processFactorial(input string) {
	//Извлекаем число из "5!" -> 5
	numStr := strings.TrimSuffix(input, "!")
	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		fmt.Println("Ошибка: неверрное число для факториала")
		return
	}
	calculateAndPrint(num, 0, "!")

}

func calculateAndPrint(a float64, b float64, operation string) {
	var result float64
	var expression string
	var validOperation bool = true

	switch operation {
	case "+":
		result = a + b
		expression = fmt.Sprintf("%.2f + %.2f = %.2f", a, b, result)
	case "-":
		result = a - b
		expression = fmt.Sprintf("%.2f - %.2f = %.2f", a, b, result)
	case "*", "x", "X", "х", "Х":
		result = a * b
		expression = fmt.Sprintf("%.2f * %.2f = %.2f", a, b, result)
	case "/":
		if b == 0 {
			fmt.Println("На ноль делить нельзя!")
			return
		}
		result = a / b
		expression = fmt.Sprintf("%.2f / %.2f = %.2f", a, b, result)

	case "^":
		result = math.Pow(a, b)
		expression = fmt.Sprintf("%.2f ^ %.2f = %.2f", a, b, result)

	case "sqrt":
		if a < 0 {
			fmt.Println("Ошибка: корень из отрицательного числа")
			return
		}
		result = math.Sqrt(a)
		expression = fmt.Sprintf("√%.2f = %.2f", a, result)

	case "%":
		result = a * b / 100
		expression = fmt.Sprintf("%.2f от %.2f%% = %.2f", a, b, result)

	case "!":
		if a < 0 || a != float64(int(a)) {
			fmt.Println("Ошибка: факторриал определен только для целых неотрицательных чисел")
			return
		}
		result = float64(factorial(int(a)))
		expression = fmt.Sprintf("%.0f! = %.0f", a, result)

	default:
		fmt.Println("Ошибка: Неизвестная операция:", operation)
		fmt.Println("Поддерживаемые операции: +, -, *, /, ^, sqrt, %, !")
		validOperation = false
	}

	if validOperation {
		fmt.Printf("Результат: %s\n", expression)
		addToHistory(expression)
	}
}

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func addToHistory(expression string) {
	history = append(history, expression)
}

func showHistory() {
	if len(history) == 0 {
		fmt.Println("История вычислений пуста.")
		return
	}

	fmt.Println("\n📊 История вычислений:")
	fmt.Println("══════════════════════════════")
	for i, expr := range history {
		fmt.Printf("%d. %s\n", i+1, expr)
	}
	fmt.Println("══════════════════════════════")
	fmt.Printf("Всего вычислений: %d\n", len(history))
}
