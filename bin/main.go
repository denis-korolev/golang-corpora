package main

import (
	"parser/bin/cmd"
	"parser/config"
)

func main() {
	config.CalculatetConfig()
	cmd.Execute()
}
