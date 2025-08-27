// history/history.go
package history

import "fmt"

// History хранит историю вычислений
var History []string

// AddToHistory добавляет выражение в историю
func AddToHistory(expression string) {
	History = append(History, expression)
}

// ShowHistory показывает историю вычислений
func ShowHistory() {
	if len(History) == 0 {
		fmt.Println("История вычислений пуста.")
		return
	}

	fmt.Println("\n📊 История вычислений:")
	fmt.Println("══════════════════════════════")
	for i, expr := range History {
		fmt.Printf("%d. %s\n", i+1, expr)
	}
	fmt.Println("══════════════════════════════")
	fmt.Printf("Всего вычислений: %d\n", len(History))
}
