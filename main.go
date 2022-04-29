package main

import (
	"LoGo/LogLevel"
	"LoGo/LoggerConfiguration"
)

func main() {
	LoggerConfiguration.
		WriteTo().
		Console().
		File("MyTestApp", "").
		MinimumLevel(LogLevel.DEBUG)
}
