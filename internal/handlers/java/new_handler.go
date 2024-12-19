package javahandler

import (
	"craft/internal/utils"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
)

type NewJavaHandler struct {
	Language    string
	BuildTool   string
	Framework   string
	TemplatesFS fs.FS
}

func (h *NewJavaHandler) SetTemplatesFS(fs fs.FS) {
	h.TemplatesFS = fs
}

const (
	placeHolder = "PROJECT_NAME"
)

func (handler *NewJavaHandler) Run(createDirectoryFor bool, projectName string) error {
	if handler.BuildTool == "" || handler.Framework == "" {
		return fmt.Errorf("invalid configuration: Build Tool or Framework not specified")
	}
	var projectDirOnHost string
	var err error

	// var constfilesThatNeedProjectNameAdjusted = []string{"docker-compose.dev.yml", "Makefile", "README.md"}
	var filesThatNeedToBeRemovedFromHostAfterwards = []string{"build.Dockerfile", "create_java_project.sh"}

	switch handler.BuildTool {
	case "maven":

		switch handler.Framework {
		case "noframework":
			projectDirOnHost, err = utils.PrepareProjectDir(createDirectoryFor, projectName)
			if err != nil {
				return fmt.Errorf("Failed to get current working directory: %v", err)
			}

			templatesPath := filepath.Join("templates", handler.Language, handler.BuildTool, handler.Framework)

			allFilesInEmbeddedFS, err := utils.ListFilesWithPattern(handler.TemplatesFS, templatesPath, "")
			fmt.Printf("all files", allFilesInEmbeddedFS)

			// 1. Copy everything over
			err = utils.CopyDirFromFS(handler.TemplatesFS, templatesPath, projectDirOnHost)
			if err != nil {
				return fmt.Errorf("Error copying over the file from the embedded folder: %v to host: %v -> error: %v\n", templatesPath, projectDirOnHost, err)
			}

			// 2. Execute ./create_java_project.sh
			scriptPath := filepath.Join(projectDirOnHost, "create_java_project.sh")

			err = os.Chmod(scriptPath, 0771)
			if err != nil {
				return fmt.Errorf("Error setting execute permissions on script: %v", err)
			}

			executeCreateJavaProjectScript := exec.Command(scriptPath, projectName)
			executeCreateJavaProjectScript.Stdout = os.Stdout
			executeCreateJavaProjectScript.Stderr = os.Stderr
			executeCreateJavaProjectScript.Dir = projectDirOnHost

			err = executeCreateJavaProjectScript.Run()
			if err != nil {
				return fmt.Errorf("Error executing script: %v", err)
			}

			// 3. remove the ./create_java_project.sh && build.Dockerfile
			for _, file := range filesThatNeedToBeRemovedFromHostAfterwards {
				filePath := filepath.Join(projectDirOnHost, file)
				err := utils.RemoveFileFromHost(filePath)
				if err != nil {
					return err
				}
			}

			// 4. if --name was passed, copy everything into the PROJECT_NAME folder
			if createDirectoryFor {
				currentPwd, err := os.Getwd()
				if err != nil {
					return err
				}

				filePaths, err := utils.GetAllFiles(currentPwd)
				if err != nil {
					return err
				}
				for _, filePath := range filePaths {
					err := utils.CopyFile(filePath, projectDirOnHost)
					if err != nil {
						return err
					}
				}
			} else {
				// 5. if -c was passed, copy everything from PROJECT_NAME folder one layer up
				// & remove PROJECT_NAME folder
				javaProjectPath := filepath.Join(projectDirOnHost, projectName)
				err := utils.CopyAllOnePathUpAndRemoveDir(javaProjectPath)
				if err != nil {
					return err
				}

			}

			// 5. change PROJECT_NAME in everyfile to projectName
			// 7. remove docker image (DOCKER_IMAGE_NAME="maven-project-generator") from host
			return fmt.Errorf("setting up a plain Java project without a framework for Maven is not yet implemented")
		case "quarkus":
			return fmt.Errorf("setting up a Quarkus project with Maven is not yet implemented")
		case "springboot":
			return fmt.Errorf("setting up a Spring Boot project with Maven is not yet implemented")
		default:
			return fmt.Errorf("unsupported framework '%s' for Maven", handler.Framework)
		}

	case "gradle":

		switch handler.Framework {
		case "noframework":
			return fmt.Errorf("setting up a plain Java project without a framework for Gradle is not yet implemented")
		case "quarkus":
			return fmt.Errorf("setting up a Quarkus project with Gradle is not yet implemented")
		case "springboot":
			return fmt.Errorf("setting up a Spring Boot project with Gradle is not yet implemented")
		default:
			return fmt.Errorf("unsupported framework '%s' for Gradle", handler.Framework)
		}

	default:
		return fmt.Errorf("unsupported build tool '%s'", handler.BuildTool)
	}
}
