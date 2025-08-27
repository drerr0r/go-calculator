// input/step_by_step.go
package input

import (
	"bufio"
	"calculator/math"
	"fmt"
	"strings"
)

// ProcessStepByStep обрабатывает пошаговый ввод
func ProcessStepByStep(reader *bufio.Reader) {
	fmt.Print("Выберите операцию:\n1 - Базовые (+, -, *, /)\n2 - Степень и корень (^, sqrt)\n3 - Проценты (%)\n4 - Факториал (!)\nВаш выбор: ")

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

	fmt.Print("Введите операцию (+, -, *, /): ")
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
	_, err = fmt.Sscanf(inputB, "%f", &b)
	if err != nil {
		fmt.Println("Ошибка чтения числа:", err)
		return
	}

	math.Calculate(a, b, operation)
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

		fmt.Print("Введите степень: ")
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

		math.Calculate(a, b, "^")

	case "2":
		var a float64
		fmt.Print("Введите число: ")
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

		math.Calculate(a, 0, "sqrt")
	}
}

func processPercentage(reader *bufio.Reader) {
	var a, b float64

	fmt.Print("Введите число: ")
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

	fmt.Print("Введите процент: ")
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

	math.Calculate(a, b, "%")
}

func processFactorialInput(reader *bufio.Reader) {
	var a float64

	fmt.Print("Введите число для вычисления факториала: ")
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

	math.Calculate(a, 0, "!")
}
