package Log

import (
	"LoGo/Color"
	"LoGo/LogLevel"
	"LoGo/LoggerConfiguration"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func Error(message string, etc ...interface{}) {
	text := fmt.Sprintf("%s [%s]: %s %v \n", time.Now().Format("2006-01-02 15:01:05"),
		LogLevel.ERROR, message, strings.Trim(fmt.Sprint(etc), "[]"))
	if LoggerConfiguration.ShouldLogToConsole() {
		fmt.Printf(Color.Red + text)
	}
	if LoggerConfiguration.ShouldLogToFile() {
		writeToFile(text)
	}
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

func writeToFile(mess string) {
	/*	message := []byte(mess)
		path := LoggerConfiguration.GetFilePath()
		err := os.WriteFile(path, message, 0644)
		if err != nil {
			return
		}*/
	path := LoggerConfiguration.GetFilePath()
	f, err := os.OpenFile(path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(mess); err != nil {
		log.Println(err)
	}
}
