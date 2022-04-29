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

func Information(message string, etc ...interface{}) {
	logEvent(0, LogLevel.INFO_MESSAGE, Color.Blue, message, etc)
}

func Warning(message string, etc ...interface{}) {
	logEvent(1, LogLevel.WARNING_MESSAGE, Color.Yellow, message, etc)
}

func Error(message string, etc ...interface{}) {
	logEvent(2, LogLevel.ERROR_MESSAGE, Color.Red, message, etc)
}

func Fatal(message string, etc ...interface{}) {
	logEvent(3, LogLevel.FATAL_MESSAGE, Color.Red, message, etc)
}

func Debug(message string, etc ...interface{}) {
	logEvent(4, LogLevel.DEBUG_MESSAGE, Color.Purple, message, etc)
}

func logEvent(minimumLevel LogLevel.MinimumLevel, levelMessage LogLevel.LogLevel, color Color.Color, message string, etc ...interface{}) {
	if LoggerConfiguration.GetMinimumLevel() >= minimumLevel {
		text := fmt.Sprintf("%s [%s]: %s %v \n", time.Now().Format("2006-01-02 15:04:05"),
			levelMessage, message, strings.Trim(fmt.Sprint(etc), "[]"))
		if LoggerConfiguration.ShouldLogToConsole() {
			fmt.Printf("%s%s", color, text)
		}
		if LoggerConfiguration.ShouldLogToFile() {
			writeToFile(text)
		}
	}
}

func writeToFile(mess string) {
	t := time.Now().Format("20060102")
	path := LoggerConfiguration.GetFilePath() + t
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
