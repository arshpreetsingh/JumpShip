package lib

import (
	"log"
	"os"
)

func StartLogs(logFile string) *log.Logger {
	// open file and create if non-existent
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	logger := log.New(file, "Custom Log", log.LstdFlags)
	return logger
}
