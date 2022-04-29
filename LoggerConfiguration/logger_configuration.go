package LoggerConfiguration

import (
	"LoGo/LogLevel"
	"os"
)

type Logger struct {
	toConsole    bool
	toDatabase   bool
	minimumLevel LogLevel.MinimumLevel
	FileLogging  FileLogging
}

type FileLogging struct {
	toFile   bool
	FilePath string
}

var myLogger = Logger{}

func WriteTo() *Logger {
	return &myLogger
}

func (l Logger) Console() *Logger {
	myLogger.toConsole = true
	return &myLogger
}

func (l Logger) File(logFileName string, path string) *Logger {
	if len(logFileName) == 0 {
		logFileName = "LoGo"
	}
	if len(path) == 0 {
		path, _ = os.UserHomeDir()
	}
	myLogger.FileLogging.toFile = true
	myLogger.FileLogging.FilePath = path + "/" + logFileName
	return &myLogger
}

func GetMinimumLevel() LogLevel.MinimumLevel {
	return myLogger.minimumLevel
}

func (l Logger) MinimumLevel(minimumLevel LogLevel.MinimumLevel) *Logger {
	myLogger.minimumLevel = minimumLevel
	return &myLogger
}

func GetFilePath() string {
	return myLogger.FileLogging.FilePath
}

func ShouldLogToFile() bool {
	return myLogger.FileLogging.toFile
}

func ShouldLogToConsole() bool {
	return myLogger.toConsole
}
