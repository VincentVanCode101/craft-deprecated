package main

import (
	"craft/cmd"
	"os"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
