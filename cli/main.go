package main

import (
	"parser/cli/cmd"
	"parser/config"
)

func main() {
	config.CalculatetConfig()
	cmd.Execute()
}
