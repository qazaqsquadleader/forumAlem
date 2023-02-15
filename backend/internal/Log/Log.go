package Log

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

func CreateLogger() (*os.File, error) {
	file, err := os.OpenFile("logger.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
	if err != nil {
		return nil, err
	}
	log.SetOutput(file)
	return file, nil
}

func LogInfo(message string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		log.Println("failed to get the runtime caller for the Logger")
	}
	log.Printf("[INFO] (file: %s, line: %d): %s\n", file, line, message)
}

func LogError(message string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		log.Println("failed to get the runtime caller for the Logger")
	}
	log.Printf("[ERROR] (file: %s, line: %d): %s\n", file, line, message)
}

func CloseLogger(logger *os.File) {
	logger.Close()
	fmt.Println("Logger closed")
}
