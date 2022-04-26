package LogLevel

type LogLevel string

const (
	ERROR   LogLevel = "ERROR"
	WARNING LogLevel = "WARNING"
	INFO    LogLevel = "INFO"
	DEBUG   LogLevel = "DEBUG"
	FATAL   LogLevel = "FATAL"
)
