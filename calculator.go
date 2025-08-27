package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Добро пожаловать в простой калькулятор на GO!")

	//Создаем читателя для ввода консоли
	reader := bufio.NewReader(os.Stdin)

	for {

		//Запрос на ввод
		fmt.Print("\nВыберрите ежим ввода:\n1 - Одна строка (например: 5 + 3)\n2 - Пошаговыйй ввод\nстоп - Выход\nВаш выбор: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		if strings.ToLower(choice) == "стоп" {
			fmt.Print("До свидания!")
			break
		}

		switch choice {
		case "1":
			processStrinleLineInput(reader)
		case "2":
			processStepByStepInput(reader)
		}
	}
}

func processStrinleLineInput(reader *bufio.Reader) {
	fmt.Print("Введите выражение (напимер 5 + 3): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	//Разбиваем строку на части
	parts := strings.Fields(input)
	if len(parts) != 3 {
		fmt.Println("Ошибка: Введите выражение в форрмате 'число оператор число'")
		fmt.Println("Пример: 5 + 3 или 12.5 * 9")
		return
	}

	//Парсим перрвое число
	a, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		fmt.Println("Ошибка: Первое значение не является числом")
		return
	}

	//Получам оператор
	operation := parts[1]

	//Парсим второе число
	b, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		fmt.Println("Ошибка: Второе значение не является числом")
		return
	}

	//Вычисляем результат
	calculateAndPrint(a, b, operation)

}

func processStepByStepInput(reader *bufio.Reader) {
	var a, b float64
	var operation string

	//Чтение первого числа
	fmt.Print("Введите первое число: ")
	inputA, _ := reader.ReadString('\n')
	inputA = strings.TrimSpace(inputA)
	_, err := fmt.Sscanf(inputA, "%f", &a)
	if err != nil {
		fmt.Println("Ошибка чтения числа:", err)
		return
	}

	//Чтение оператора
	fmt.Print("Введите операцию (+, -, *, /): ")
	inputOp, _ := reader.ReadString('\n')
	operation = strings.TrimSpace(inputOp)

	//Чтение второго числа
	fmt.Print("Введите второго число: ")
	inputB, _ := reader.ReadString('\n')
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
	default:
		fmt.Println("Ошибка: Неизвесстная операция:", operation)
		fmt.Println("Поддерживаемые опеации: +, -, *, / ")
		validOperation = false

	}

	if validOperation {
		fmt.Printf("Результат: %.2f %s %.2f = %.2f\n\n", a, operation, b, result)

	}
}
