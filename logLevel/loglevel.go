package logLevel

type LogLevel string

const (
	ERROR       LogLevel = "ERROR"
	WARNING     LogLevel = "WARNING"
	INFORMATION LogLevel = "INFORMATION"
	FATAL       LogLevel = "FATAL"
)
