package LogLevel

type LogLevel string
type MinimumLevel int

const (
	ERROR_MESSAGE   LogLevel = "ERROR"
	WARNING_MESSAGE LogLevel = "WARNING"
	INFO_MESSAGE    LogLevel = "INFO"
	DEBUG_MESSAGE   LogLevel = "DEBUG"
	FATAL_MESSAGE   LogLevel = "FATAL"
)

const (
	INFO MinimumLevel = iota
	WARNING
	ERROR
	FATAL
	DEBUG
)
