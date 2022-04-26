package LoggerConfiguration

type Logger struct {
	msg string
}

func WriteTo() *Logger {
	return &Logger{msg: "test"}
}

func (l *Logger) Console() {
	print("Write to console")
}

func (l *Logger) File() {
	print("Write to file")
}
