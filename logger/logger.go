package logger

import (
	"fmt"
	"log"
	"os"
)

// Logger инкапсулирует логирование в файл и консоль.
type Logger struct {
	*log.Logger
}

// NewLogger создает новый экземпляр Logger. Логи могут записываться как в файл, так и в консоль.
func NewLogger(logFile string) (*Logger, error) {
	// Открытие файла для записи логов, если файл не существует, он будет создан.
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл для логирования: %v", err)
	}

	// Создание логера, который будет писать как в файл, так и в консоль.
	logger := log.New(file, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)

	return &Logger{logger}, nil
}
