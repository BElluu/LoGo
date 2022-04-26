package Log

import (
	"LoGo/logLevel"
	"fmt"
	"time"
)

func Error(message string) {
	fmt.Printf("%s [%s]: %s", time.Now().Format("2006-01-02 15:01:05"), logLevel.ERROR, message)
}

func Warning(message string) {
	fmt.Printf("%s [%s]: %s", time.Now().Format("2006-01-02 15:01:05"), logLevel.WARNING, message)
}

func Information(message string) {
	fmt.Printf("%s [%s]: %s", time.Now().Format("2006-01-02 15:01:05"), logLevel.INFORMATION, message)
}

func Fatal(message string) {
	fmt.Printf("%s [%s]: %s", time.Now().Format("2006-01-02 15:01:05"), logLevel.FATAL, message)
}
