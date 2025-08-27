package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Добро пожаловать в простой калькулятор на GO!")

	//Создаем читателя для ввода консоли
	reader := bufio.NewReader(os.Stdin)

	for {
		var a, b float64
		var operation string

		//Запрос на ввод
		fmt.Println("Ведите 'стоп' для выхода или любой другой символ для продолжения: ")
		exitFlag, _ := reader.ReadString('\n')
		exitFlag = strings.TrimSpace(exitFlag)

		if strings.ToLower(exitFlag) == "стоп" {
			fmt.Print("До свидания!")
			break
		}

		//Чтение первого числа
		fmt.Print("Введите первое число: ")
		inputA, _ := reader.ReadString('\n')
		inputA = strings.TrimSpace(inputA)
		_, err := fmt.Sscanf(inputA, "%f", &a)
		if err != nil {
			fmt.Println("Ошибка чтения числа:", err)
			continue
		}

		//Чтение оператора
		fmt.Print("Введите операцию (+, -, *, /): ")
		inputOp, _ := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Ошибка чтения опперации:", err)
			continue
		}
		operation = strings.TrimSpace(inputOp)

		//Чтение второго числа
		fmt.Print("Введите второго число: ")
		inputB, _ := reader.ReadString('\n')
		inputB = strings.TrimSpace(inputB)
		_, err = fmt.Sscanf(inputB, "%f", &b)
		if err != nil {
			fmt.Println("Ошибка чтения числа:", err)
			continue
		}

		var result float64

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
				continue
			}
			result = a / b
		default:
			fmt.Println("Ошибка: Неизвесстная операция:", operation)
			continue
		}

		fmt.Printf("Результат: %.2f %s %.2f = %.2f\n\n", a, operation, b, result)

	}

}
