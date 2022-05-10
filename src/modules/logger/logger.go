package logger

import (
	"api/src/config"
	"fmt"
	"os"
	"time"
)

func Debug(message string) {
	createMessage("DEBUG", message)
}

func Info(message string) {
	createMessage("INFO", message)
}

func Warning(message string) {
	createMessage("WARNING", message)
}

func Error(message string) {
	createMessage("ERROR", message)
	os.Exit(config.DB_ERROR)
}

func createMessage(log_type string, message string) {
	now := time.Now()
	fmt.Println(log_type, " | ", now.Format("02.01.2006 15:04:05"), " | ", message)
}

func write() {}
