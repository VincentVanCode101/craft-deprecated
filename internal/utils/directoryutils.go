package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// PrepareProjectDir prepares a project directory.
// If `createDir` is true, it creates a new directory for the project.
// If `createDir` is false, it returns the current working directory.
// `projectName` is used to name the new directory if created.
func PrepareProjectDir(createDir bool, projectName string) (string, error) {
	baseDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("could not get current directory: %w", err)
	}

	if createDir {
		projectDir := filepath.Join(baseDir, projectName)
		if err := os.Mkdir(projectDir, 0755); err != nil {
			return "", fmt.Errorf("could not create project directory: %w", err)
		}
		return projectDir, nil
	}

	return baseDir, nil
}
