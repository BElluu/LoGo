package LoggerConfiguration

type Logger struct {
	ToConsole bool
	ToFile    bool
}

func WriteTo() Logger {
	return Logger{}
}

func (l Logger) Console() Logger {
	l.ToConsole = true
	print("Write to console")
	return l
}

func (l Logger) File(path string) {
	l.ToFile = true
	print("Write to file %s", path)
}
