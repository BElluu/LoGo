package Log

import (
	"LoGo/Color"
	"LoGo/LogLevel"
	"fmt"
	"strings"
	"time"
)

func Error(message string, etc ...interface{}) {
	fmt.Printf(Color.Red+"%s [%s]: %s %v \n", time.Now().Format("2006-01-02 15:01:05"),
		LogLevel.ERROR, message, strings.Trim(fmt.Sprint(etc), "[]"))
}

func Warning(message string, etc ...interface{}) {
	fmt.Printf(Color.Yellow+"%s [%s]: %s %v \n", time.Now().Format("2006-01-02 15:01:05"),
		LogLevel.WARNING, message, strings.Trim(fmt.Sprint(etc), "[]"))
}

func Information(message string, etc ...interface{}) {
	fmt.Printf(Color.Blue+"%s [%s]: %s %v \n", time.Now().Format("2006-01-02 15:01:05"),
		LogLevel.INFO, message, strings.Trim(fmt.Sprint(etc), "[]"))
}

func Debug(message string, etc ...interface{}) {
	fmt.Printf(Color.Purple+"%s [%s]: %s %v \n", time.Now().Format("2006-01-02 15:01:05"),
		LogLevel.DEBUG, message, strings.Trim(fmt.Sprint(etc), "[]"))
}

func Fatal(message string, etc ...interface{}) {
	fmt.Printf(Color.Red+"%s [%s]: %s %v \n", time.Now().Format("2006-01-02 15:01:05"),
		LogLevel.FATAL, message, strings.Trim(fmt.Sprint(etc), "[]"))
}
