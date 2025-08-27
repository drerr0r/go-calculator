// math/operations.go
package math

import (
	"calculator/history"
	"fmt"
	"math"
)

// Calculate выполняет математические операции
func Calculate(a float64, b float64, operation string) {
	var result float64
	var expression string

	switch operation {
	case "+":
		result = a + b
		expression = fmt.Sprintf("%.2f + %.2f = %.2f", a, b, result)
	case "-":
		result = a - b
		expression = fmt.Sprintf("%.2f - %.2f = %.2f", a, b, result)
	case "*":
		result = a * b
		expression = fmt.Sprintf("%.2f * %.2f = %.2f", a, b, result)
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль!")
			return
		}
		result = a / b
		expression = fmt.Sprintf("%.2f / %.2f = %.2f", a, b, result)
	case "^":
		result = math.Pow(a, b)
		expression = fmt.Sprintf("%.2f ^ %.2f = %.2f", a, b, result)
	case "sqrt":
		if a < 0 {
			fmt.Println("Ошибка: корень из отрицательного числа!")
			return
		}
		result = math.Sqrt(a)
		expression = fmt.Sprintf("√%.2f = %.2f", a, result)
	case "%":
		result = a * b / 100
		expression = fmt.Sprintf("%.2f от %.2f%% = %.2f", a, b, result)
	case "!":
		if a < 0 || a != float64(int(a)) {
			fmt.Println("Ошибка: факториал только для целых неотрицательных чисел")
			return
		}
		result = float64(factorial(int(a)))
		expression = fmt.Sprintf("%.0f! = %.0f", a, result)
	default:
		fmt.Println("Ошибка: Неизвестная операция:", operation)
		return
	}

	fmt.Printf("Результат: %s\n", expression)
	history.AddToHistory(expression)
}

// factorial вычисляет факториал
func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
