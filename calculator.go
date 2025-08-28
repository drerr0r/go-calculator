// main.go
package main

import (
	"bufio"
	"calculator/history"
	"calculator/input"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Добро пожаловать в продвинутый калькулятор на Go!")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nВыберите режим ввода:\n1 - Одна строка\n2 - Пошаговый ввод\nистория - История вычислений\nстоп - Выход\nВаш выбор: ")

		choice, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения ввода:", err)
			continue
		}

		choice = strings.TrimSpace(choice)

		switch strings.ToLower(choice) {
		case "стоп":
			fmt.Println("До свидания!")
			return
		case "история":
			history.ShowHistory()
		case "1":
			input.ProcessSingleLine(reader)
		case "2":
			input.ProcessStepByStep(reader)
		default:
			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
	}
}
