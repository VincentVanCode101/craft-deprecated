package gohandler

import (
	"fmt"
	"os"
	"path/filepath"
)

type NewGoHandler struct {
	Language string
}

func (h NewGoHandler) Run(createDirectoryFor bool, projectName string) error {
	fmt.Printf("Creating a new Go project with name '%s'...\n", projectName)

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return err
	}

	if createDirectoryFor {

		err = os.Mkdir(projectName, 0755)

		if err != nil {
			fmt.Println("Error creating directory: %v", projectName, err)
			return err
		}
	}

	goTemplatesPath := filepath.Join("/templates", h.Language)
	err = os.Chdir(goTemplatesPath)
	if err != nil {
		fmt.Println("Error changing to /templates/go:", err)
		return err
	}

	fmt.Println("Current Directory:", currentDir)

	entries, err := os.ReadDir(currentDir)

	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Println("[DIR] ", entry.Name())
		} else {
			fmt.Println("[FILE]", entry.Name())
		}
	}
	return nil
}

type ScanGoHandler struct{}

func (h ScanGoHandler) Run() error {
	fmt.Println("Scanning a Go project...")
	return nil
}
