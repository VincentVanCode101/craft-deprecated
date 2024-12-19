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

const placeHolder = "PROJECT_NAME"

func (handler *NewJavaHandler) Run(createDirectoryFor bool, projectName string) error {
	if handler.BuildTool == "" || handler.Framework == "" {
		return fmt.Errorf("invalid configuration: Build Tool or Framework not specified")
	}

	projectDirOnHost, err := utils.PrepareProjectDir(createDirectoryFor, projectName)
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %v", err)
	}

	switch handler.BuildTool {
	case "maven":
		return handler.handleMavenProject(projectDirOnHost, projectName)
	case "gradle":
		return handler.handleGradleProject()
	default:
		return fmt.Errorf("unsupported build tool '%s'", handler.BuildTool)
	}
}

func (handler *NewJavaHandler) handleMavenProject(projectDirOnHost, projectName string) error {
	switch handler.Framework {
	case "noframework":
		return handler.setupNoFrameworkMavenProject(projectDirOnHost, projectName)
	case "quarkus":
		return fmt.Errorf("setting up a Quarkus project with Maven is not yet implemented")
	case "springboot":
		return fmt.Errorf("setting up a Spring Boot project with Maven is not yet implemented")
	default:
		return fmt.Errorf("unsupported framework '%s' for Maven", handler.Framework)
	}
}

func (handler *NewJavaHandler) handleGradleProject() error {
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
}

func (handler *NewJavaHandler) setupNoFrameworkMavenProject(projectDirOnHost, projectName string) error {
	filesThatNeedProjectNameAdjustedOnce := []string{"docker-compose.dev.yml", "Makefile", "README.md"}
	filesThatNeedProjectNameAdjustedEverywhere := []string{"README.md"}
	filesThatNeedToBeRemoved := []string{"build.Dockerfile", "create_java_project.sh"}

	templatesPath := filepath.Join("templates", handler.Language, handler.BuildTool, handler.Framework)
	if err := handler.copyTemplateFilesToHost(templatesPath, projectDirOnHost); err != nil {
		return err
	}

	scriptPath := filepath.Join(projectDirOnHost, "create_java_project.sh")
	if err := handler.executeProjectSetupScript(scriptPath, projectName, projectDirOnHost); err != nil {
		return err
	}

	if err := handler.cleanupFiles(projectDirOnHost, filesThatNeedToBeRemoved); err != nil {
		return err
	}

	javaProjectPath := filepath.Join(projectDirOnHost, projectName)
	if err := utils.CopyAllOnePathUpAndRemoveDir(javaProjectPath); err != nil {
		return err
	}

	if err := handler.adjustProjectNames(projectDirOnHost, filesThatNeedProjectNameAdjustedOnce, filesThatNeedProjectNameAdjustedEverywhere, projectName); err != nil {
		return err
	}

	return fmt.Errorf("setting up a plain Java project without a framework for Maven is not yet implemented")
}

func (handler *NewJavaHandler) copyTemplateFilesToHost(templatesPath, projectDirOnHost string) error {
	if err := utils.CopyDirFromFS(handler.TemplatesFS, templatesPath, projectDirOnHost); err != nil {
		return fmt.Errorf("error copying files from template path: %v", err)
	}
	return nil
}

func (handler *NewJavaHandler) executeProjectSetupScript(scriptPath, projectName, projectDirOnHost string) error {
	if err := os.Chmod(scriptPath, 0771); err != nil {
		return fmt.Errorf("error setting execute permissions on script: %v", err)
	}

	execCmd := exec.Command(scriptPath, projectName)
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr
	execCmd.Dir = projectDirOnHost

	if err := execCmd.Run(); err != nil {
		return fmt.Errorf("error executing setup script: %v", err)
	}
	return nil
}

func (handler *NewJavaHandler) cleanupFiles(projectDirOnHost string, files []string) error {
	for _, file := range files {
		filePath := filepath.Join(projectDirOnHost, file)
		if err := utils.RemoveFileFromHost(filePath); err != nil {
			return fmt.Errorf("error removing file '%s': %v", filePath, err)
		}
	}
	return nil
}

func (handler *NewJavaHandler) adjustProjectNames(projectDirOnHost string, onceFiles, everywhereFiles []string, projectName string) error {
	for _, filePath := range onceFiles {
		if err := utils.ChangeWordInFile(filepath.Join(projectDirOnHost, filePath), placeHolder, projectName, false); err != nil {
			return fmt.Errorf("error adjusting project name in file '%s': %v", filePath, err)
		}
	}

	for _, filePath := range everywhereFiles {
		if err := utils.ChangeWordInFile(filepath.Join(projectDirOnHost, filePath), placeHolder, projectName, true); err != nil {
			return fmt.Errorf("error adjusting project name in file '%s': %v", filePath, err)
		}
	}
	return nil
}
