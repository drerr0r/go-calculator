// math/operations.go
package math

import (
	"calculator/history"
	"fmt"
)

// Calculate выполняет математические операции с использованием парсера
func Calculate(expression string) (float64, error) {
	parser := NewParser(expression)
	result, err := parser.Parse()
	if err != nil {
		return 0, err
	}

	// Сохраняем в историю
	historyEntry := fmt.Sprintf("%s = %.2f", expression, result)
	if err := history.Add(historyEntry); err != nil {
		return 0, err
	}

	return result, nil
}

// CalculateLegacy старая функция для обратной совместимости
func CalculateLegacy(a float64, b float64, operation string) {
	var result float64
	var expr string

	switch operation {
	case "+":
		result = a + b
		expr = fmt.Sprintf("%.2f + %.2f", a, b)
	case "-":
		result = a - b
		expr = fmt.Sprintf("%.2f - %.2f", a, b)
	case "*":
		result = a * b
		expr = fmt.Sprintf("%.2f * %.2f", a, b)
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль!")
			return
		}
		result = a / b
		expr = fmt.Sprintf("%.2f / %.2f", a, b)
	case "^":
		result = Pow(a, b)
		expr = fmt.Sprintf("%.2f ^ %.2f", a, b)
	case "sqrt":
		if a < 0 {
			fmt.Println("Ошибка: корень из отрицательного числа!")
			return
		}
		result = Sqrt(a)
		expr = fmt.Sprintf("√%.2f", a)
	case "%":
		result = a * b / 100
		expr = fmt.Sprintf("%.2f от %.2f%%", a, b)
	case "!":
		if a < 0 || a != float64(int(a)) {
			fmt.Println("Ошибка: факториал только для целых неотрицательных чисел")
			return
		}
		result = float64(factorial(int(a)))
		expr = fmt.Sprintf("%.0f!", a)
	default:
		fmt.Println("Ошибка: Неизвестная операция:", operation)
		return
	}

	fmt.Printf("Результат: %s = %.2f\n", expr, result)
	history.Add(fmt.Sprintf("%s = %.2f", expr, result))
}

// Sqrt вычисляет квадратный корень
func Sqrt(num float64) float64 {
	// Простая реализация методом Ньютона
	if num == 0 {
		return 0
	}
	result := num
	for i := 0; i < 10; i++ {
		result = (result + num/result) / 2
	}
	return result
}

// factorial вычисляет факториал
func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
