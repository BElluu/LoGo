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

// TODO Get minimum level in variable. Do not call method in if statement

func Information(message string, etc ...interface{}) {
	text := fmt.Sprintf("%s [%s]: %s %v \n", time.Now().Format("2006-01-02 15:01:05"),
		LogLevel.INFO_MESSAGE, message, strings.Trim(fmt.Sprint(etc), "[]"))
	if LoggerConfiguration.ShouldLogToConsole() {
		fmt.Printf(Color.Blue + text)
	}
	if LoggerConfiguration.ShouldLogToFile() {
		writeToFile(text)
	}
}

func Warning(message string, etc ...interface{}) {
	if LoggerConfiguration.GetMinimumLevel() >= 1 {
		text := fmt.Sprintf("%s [%s]: %s %v \n", time.Now().Format("2006-01-02 15:01:05"),
			LogLevel.WARNING_MESSAGE, message, strings.Trim(fmt.Sprint(etc), "[]"))
		if LoggerConfiguration.ShouldLogToConsole() {
			fmt.Printf(Color.Yellow + text)
		}
		if LoggerConfiguration.ShouldLogToFile() {
			writeToFile(text)
		}
	}
}

func Error(message string, etc ...interface{}) {
	if LoggerConfiguration.GetMinimumLevel() >= 2 {
		text := fmt.Sprintf("%s [%s]: %s %v \n", time.Now().Format("2006-01-02 15:01:05"),
			LogLevel.ERROR_MESSAGE, message, strings.Trim(fmt.Sprint(etc), "[]"))
		if LoggerConfiguration.ShouldLogToConsole() {
			fmt.Printf(Color.Red + text)
		}
		if LoggerConfiguration.ShouldLogToFile() {
			writeToFile(text)
		}
	}
}

func Fatal(message string, etc ...interface{}) {
	if LoggerConfiguration.GetMinimumLevel() >= 3 {
		text := fmt.Sprintf("%s [%s]: %s %v \n", time.Now().Format("2006-01-02 15:01:05"),
			LogLevel.FATAL_MESSAGE, message, strings.Trim(fmt.Sprint(etc), "[]"))
		if LoggerConfiguration.ShouldLogToConsole() {
			fmt.Printf(Color.Red + text)
		}
		if LoggerConfiguration.ShouldLogToFile() {
			writeToFile(text)
		}
	}
}

func Debug(message string, etc ...interface{}) {
	if LoggerConfiguration.GetMinimumLevel() >= 4 {
		text := fmt.Sprintf("%s [%s]: %s %v \n", time.Now().Format("2006-01-02 15:01:05"),
			LogLevel.DEBUG_MESSAGE, message, strings.Trim(fmt.Sprint(etc), "[]"))
		if LoggerConfiguration.ShouldLogToConsole() {
			fmt.Printf(Color.Purple + text)
		}
		if LoggerConfiguration.ShouldLogToFile() {
			writeToFile(text)
		}
	}
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
