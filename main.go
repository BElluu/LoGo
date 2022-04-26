package main

import (
	"LoGo/Log"
	"LoGo/LoggerConfiguration"
)

func main() {
	Log.Error("Panikujemy")
	LoggerConfiguration.WriteTo().Console()
}
