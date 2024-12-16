package gohandler

import (
	"fmt"
	"os"
	"path/filepath"

	"craft/internal/utils"
)

type NewGoHandler struct {
	Language string
}

var (
	dockerComposeFile                = "docker-compose.dev.yml"
	filesThatNeedProjectNameAdjusted = []string{"go.mod", dockerComposeFile, "Makefile"}
	filesToCopyOver                  = append(filesThatNeedProjectNameAdjusted, "Dockerfile", "go.sum", "pre-commit")
)

func (h NewGoHandler) Run(createDirectoryFor bool, projectName string) error {
	goTemplatesPathInContainer := filepath.Join("/templates", h.Language)

	fmt.Printf("Creating a new Go project with name '%s'...\n", projectName)

	projectDir, err := utils.PrepareProjectDir(createDirectoryFor, projectName)
	if err != nil {
		return fmt.Errorf("failed to prepare project directory: %w", err)
	}

	err = os.Chdir(goTemplatesPathInContainer)
	if err != nil {
		return fmt.Errorf("error changing to /templates/go: %w", err)
	}

	for _, file := range filesToCopyOver {
		filePathOnHost := utils.GetFilePath(projectDir, file)
		err = utils.CopyFile(file, filePathOnHost)
		if err != nil {
			fmt.Printf("Error copying file: %v\n", err)
		} else {
			fmt.Printf("File copied successfully from %s to %s\n", file, filePathOnHost)
		}
	}

	err = os.Chdir(projectDir)
	if err != nil {
		return fmt.Errorf("error changing to project directory %s: %w", projectDir, err)
	}

	for _, filePath := range filesThatNeedProjectNameAdjusted {
		err := utils.ChangeProjectNameInFile(filePath, projectName)
		if err != nil {
			fmt.Printf("Error changing the project name in %s: %v\n", filePath, err)
		}
	}

	return nil
}
