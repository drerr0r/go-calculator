package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var history []string

func main() {
	fmt.Println("Добро пожаловать в простой калькулятор на Go!")

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
	fmt.Print("Введите выражение (например 5 + 3): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения ввода:", err)
		return
	}
	input = strings.TrimSpace(input)

	// Обрабатываем оба формата: "5+2" и "5 + 2"
	processedInput := addSpacesAroundOperators(input)

	// Разбиваем строку на части
	parts := strings.Fields(processedInput)
	if len(parts) != 3 {
		fmt.Println("Ошибка: Введите выражение в формате 'число оператор число'")
		fmt.Println("Пример: 5 + 3, 10.5*2, 7-4, 15/3")
		return
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
	operators := []string{"+", "-", "/", "*", "x", "X", "х", "Х"}

	for _, op := range operators {
		input = strings.ReplaceAll(input, op, " "+op+" ")
	}
	return input
}

func processStepByStepInput(reader *bufio.Reader) {
	var a, b float64
	var operation string

	// Чтение первого числа
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

	// Чтение оператора
	fmt.Print("Введите операцию (+, -, *, /, х): ")
	inputOp, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения ввода:", err)
		return
	}
	operation = strings.TrimSpace(inputOp)

	// Чтение второго числа
	fmt.Print("Введите второе число: ")
	inputB, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения ввода:", err)
		return
	}
	inputB = strings.TrimSpace(inputB)
	_, err = fmt.Sscanf(inputB, "%f", &b)
	if err != nil {
		fmt.Println("Ошибка чтения числа:", err)
		return
	}
	calculateAndPrint(a, b, operation)
}

func calculateAndPrint(a float64, b float64, operation string) {
	var result float64
	var validOperation bool = true

	switch operation {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("На ноль делить нельзя!")
			return
		}
		result = a / b
	case "x", "X", "х", "Х":
		result = a * b
	default:
		fmt.Println("Ошибка: Неизвестная операция:", operation)
		fmt.Println("Поддерживаемые операции: +, -, *, /")
		validOperation = false
	}

	if validOperation {
		// ← ДОБАВЛЕНО: создаем expression
		expression := fmt.Sprintf("%.2f %s %.2f = %.2f", a, operation, b, result)
		fmt.Printf("Результат: %s\n\n", expression)

		// Добавляем выражение в историю
		addToHistory(expression)
	}
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
