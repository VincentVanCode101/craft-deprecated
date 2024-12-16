package main

import (
	"craft/cmd"
	"embed"
	"os"
)

// Embed the entire templates directory.
//
//go:embed templates/*
var templatesFS embed.FS

func main() {
	err := cmd.Execute(templatesFS)
	if err != nil {
		os.Exit(1)
	}
}
