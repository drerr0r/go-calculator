// calculator.go
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
	// Загружаем историю при запуске
	err := history.Init()
	if err != nil {
		fmt.Printf("Ошибка загрузки истории: %v\n", err)
	}

	defer func() {
		// Сохраняем историю при выходе
		if err := history.Save(); err != nil {
			fmt.Printf("Ошибка сохранения истории: %v\n", err)
		}
	}()

	fmt.Println("Добро пожаловать в продвинутый калькулятор на Go!")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nВыберите режим ввода:\n1 - Одна строка\n2 - Пошаговый ввод\n3 - История вычислений\n4 - Очистить историю\n5 - Выход\nВаш выбор: ")

		choice, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения ввода:", err)
			continue
		}

		choice = strings.TrimSpace(choice)

		switch choice {
		case "5":
			fmt.Println("До свидания!")
			return
		case "3":
			showHistory()
		case "4":
			clearHistory()
		case "1":
			input.ProcessSingleLine(reader)
		case "2":
			input.ProcessStepByStep(reader)
		default:
			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
	}
}

func showHistory() {
	entries := history.Get()
	if len(entries) == 0 {
		fmt.Println("История пуста")
		return
	}

	fmt.Println("\nИстория операций:")
	for i, entry := range entries {
		fmt.Printf("%d. %s\n", i+1, entry)
	}
}

func clearHistory() {
	if err := history.Clear(); err != nil {
		fmt.Printf("Ошибка очистки истории: %v\n", err)
	} else {
		fmt.Println("История очищена")
	}
}
