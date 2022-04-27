package main

import (
	"LoGo/Log"
	"LoGo/LogLevel"
	"LoGo/LoggerConfiguration"
)

func main() {
	LoggerConfiguration.
		WriteTo().
		Console().
		File("/home/bartek/testlog/log.txt", 1).
		MinimumLevel(LogLevel.DEBUG)
	wtf := LoggerConfiguration.GetMinimumLevel()
	println(wtf)
	Log.Error("Error! Something is wrong.", "Host:Localhost", "Port:1123", "Moreeeeexd")
	Log.Error("5", 5, true, byte(131))
	Log.Debug("Dev time")
	Log.Warning("Watch out! Warning...")
	Log.Information("Just information")
	Log.Fatal("Critical ERROR")
}
