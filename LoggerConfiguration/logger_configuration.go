package LoggerConfiguration

type Logger struct {
	toConsole  bool
	toFile     bool
	toDatabase bool
	FilePath   string
}

var myLogger = Logger{}

func WriteTo() *Logger {
	return &myLogger
}

func (l Logger) Console() *Logger {
	myLogger.toConsole = true
	return &myLogger
}

func (l Logger) File(path string) *Logger {
	myLogger.toFile = true
	myLogger.FilePath = path
	return &myLogger
}

func GetFilePath() string {
	return myLogger.FilePath
}

func ShouldLogToFile() bool {
	return myLogger.toFile
}

func ShouldLogToConsole() bool {
	return myLogger.toConsole
}
