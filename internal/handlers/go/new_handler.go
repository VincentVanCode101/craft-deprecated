package gohandler

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"craft/internal/utils"
)

type NewGoHandler struct {
	Language    string
	TemplatesFS fs.FS
}

func (h *NewGoHandler) SetTemplatesFS(fs fs.FS) {
	h.TemplatesFS = fs
}

var (
	filesThatNeedProjectNameAdjusted = []string{"go.mod.template", "docker-compose.dev.yml", "Makefile"}
	filesToCopyOver                  = append(filesThatNeedProjectNameAdjusted, "Dockerfile", "go.sum.template", "pre-commit")
)

func (handler *NewGoHandler) Run(createDirectoryFor bool, projectName string) error {
	var projectDir string
	var err error

	if utils.IsRunningInDocker() {
		fmt.Println("Running inside a Docker container.")

		projectDir, err = utils.PrepareProjectDir(createDirectoryFor, projectName)
		if err != nil {
			fmt.Printf("Failed to prepare project directory: %w\n", err)
			return err
		}

		err = os.Chdir(filepath.Join("/templates", handler.Language))
		if err != nil {
			fmt.Printf("Error changing directory to templates: %v\n", err)
			return err
		}
	} else {
		fmt.Println("Not running inside a Docker container.")
		// Use the binary's working directory
		projectDir, err = utils.PrepareProjectDir(createDirectoryFor, projectName)
		if err != nil {
			fmt.Printf("Failed to get current working directory: %v\n", err)
			return err
		}
	}

	filesThatHaveBeenCopiedOver := make([]string, 0)

	for _, file := range filesToCopyOver {
		var fileContent []byte

		if utils.IsRunningInDocker() {
			// Read file content directly from the container's filesystem
			filePath := filepath.Join("/templates", handler.Language, file)
			fileContent, err = os.ReadFile(filePath)
			if err != nil {
				fmt.Printf("Error reading file %s: %v\n", filePath, err)
				continue
			}
		} else {
			// Read file content from the embedded filesystem
			embeddedFilePath := filepath.Join("templates", handler.Language, file)
			fileContent, err = fs.ReadFile(handler.TemplatesFS, embeddedFilePath)
			if err != nil {
				fmt.Printf("Error reading embedded file %s: %v\n", embeddedFilePath, err)
				continue
			}
		}

		// Write the file content to the project directory
		filePathOnHost := utils.GetFilePath(projectDir, file)
		err = os.WriteFile(filePathOnHost, fileContent, 0644)
		if err != nil {
			fmt.Printf("Error writing file %s: %v\n", filePathOnHost, err)
			continue
		}

		fmt.Printf("File copied successfully from %s to %s\n", file, filePathOnHost)

		// Remove the .template suffix if present
		if strings.HasSuffix(file, ".template") {
			newFileName := strings.TrimSuffix(filePathOnHost, ".template")
			err = os.Rename(filePathOnHost, newFileName)
			if err != nil {
				fmt.Printf("Error renaming file %s to %s: %v\n", filePathOnHost, newFileName, err)
			} else {
				fmt.Printf("Renamed file %s to %s\n", filePathOnHost, newFileName)
				filesThatHaveBeenCopiedOver = append(filesThatHaveBeenCopiedOver, filepath.Base(newFileName)) // Add adjusted file to list
			}
		} else {
			filesThatHaveBeenCopiedOver = append(filesThatHaveBeenCopiedOver, filepath.Base(file)) // Add original name if not adjusted
		}
	}

	// Change directory to the project directory
	err = os.Chdir(projectDir)
	if err != nil {
		fmt.Printf("Error changing to project directory %s: %w\n", projectDir, err)
		return err
	}

	// Normalize filesThatNeedProjectNameAdjusted by stripping ".template"
	normalizedFilesThatNeedProjectNameAdjusted := make(map[string]bool)
	for _, file := range filesThatNeedProjectNameAdjusted {
		normalizedFilesThatNeedProjectNameAdjusted[strings.TrimSuffix(file, ".template")] = true
	}

	// Adjust project name in copied files
	for _, filePath := range filesThatHaveBeenCopiedOver {
		// Check if the file is in the normalized list of files needing project name adjustment
		if normalizedFilesThatNeedProjectNameAdjusted[filePath] {
			err := utils.ChangeProjectNameInFile(filePath, projectName)
			if err != nil {
				fmt.Printf("Error changing the project name in %s: %v\n", filePath, err)
			}
		}
	}

	return nil
}
