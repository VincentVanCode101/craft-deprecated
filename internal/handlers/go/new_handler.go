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
	dockerComposeFile                = "docker-compose.dev.yml"
	filesThatNeedProjectNameAdjusted = []string{"go.mod.template", dockerComposeFile, "Makefile"}
	filesToCopyOver                  = append(filesThatNeedProjectNameAdjusted, "Dockerfile", "go.sum.template", "pre-commit")
)

// Check if running in Docker by looking for /.dockerenv or entries in /proc/1/cgroup
func isRunningInDocker() bool {
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}

	file, err := os.Open("/proc/1/cgroup")
	if err != nil {
		return false
	}
	defer file.Close()

	buf := make([]byte, 4096)
	n, _ := file.Read(buf)
	return strings.Contains(string(buf[:n]), "docker")
}

func (handler *NewGoHandler) Run(createDirectoryFor bool, projectName string) error {
	var projectDir string
	var err error

	if isRunningInDocker() {
		fmt.Println("Running inside a Docker container.")
		// Use utils.PrepareProjectDir to determine the directory
		projectDir, err = utils.PrepareProjectDir(createDirectoryFor, projectName)
		if err != nil {
			fmt.Printf("Failed to prepare project directory: %w\n", err)
			return err
		}

		// Set the source directory for templates within the container
		err = os.Chdir(filepath.Join("/templates", handler.Language))
		if err != nil {
			fmt.Printf("Error changing directory to templates: %v\n", err)
			return err
		}
	} else {
		fmt.Println("Not running inside a Docker container.")
		// Use the binary's working directory
		projectDir, err = os.Getwd()
		if err != nil {
			fmt.Printf("Failed to get current working directory: %v\n", err)
			return err
		}
	}

	// Track renamed files
	filesThatHaveBeenCopiedOver := make([]string, 0)

	for _, file := range filesToCopyOver {
		var fileContent []byte

		if isRunningInDocker() {
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

	// Adjust project name in copied files
	for _, filePath := range filesThatHaveBeenCopiedOver {
		err := utils.ChangeProjectNameInFile(filePath, projectName)
		if err != nil {
			fmt.Printf("Error changing the project name in %s: %v\n", filePath, err)
		}
	}

	// time.Sleep(60 * time.Second)
	return nil
}
