// history/history.go
package history

import (
	"encoding/json"
	"os"
	"sync"
)

var (
	history     []string
	historyFile = "calculator_history.json"
	mu          sync.Mutex
)

// Init загружает историю из файла при запуске
func Init() error {
	mu.Lock()
	defer mu.Unlock()

	file, err := os.Open(historyFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	decorder := json.NewDecoder(file)
	return decorder.Decode(&history)
}

// Save сохраняет историю в файл
func Save() error {
	mu.Lock()
	defer mu.Unlock()

	file, err := os.Create(historyFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(history)

}

// Add добавляет запись в историю и автоматически сохраняет
func Add(entry string) error {
	mu.Lock()
	defer mu.Unlock() // Должно быть здесь

	history = append(history, entry)
	return Save()
}

// Get возвращает всю историю
func Get() []string {
	mu.Lock()
	defer mu.Unlock()
	return append([]string{}, history...)

}

// Clear очищает историю и удаляет файл
func Clear() error {
	mu.Lock()
	defer mu.Unlock() // Должно быть здесь

	history = []string{}
	return os.Remove(historyFile)
}

// GetLast возвращает последнюю запись в истории
func GetLast() string {
	mu.Lock()
	defer mu.Unlock()

	if len(history) == 0 {
		return ""
	}
	return history[len(history)-1]
}
