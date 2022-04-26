package main

import (
	"LoGo/Log"
	"LoGo/LoggerConfiguration"
)

func main() {
	LoggerConfiguration.WriteTo().File("/home/bartek/testlog/log.txt").Console()
	Log.Error("Error! Something is wrong.", "Host:Localhost", "Port:1123", "Moreeeeexd")
	/*	Log.Error("5", 5, true, byte(131))
		Log.Debug("Dev time")
		Log.Warning("Watch out! Warning...")
		Log.Information("Just information")
		Log.Fatal("Critical ERROR")*/
}
