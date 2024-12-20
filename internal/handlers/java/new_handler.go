package javahandler

import (
	"craft/internal/utils"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type NewJavaHandler struct {
	Dependencies []string
	Language     string
	BuildTool    string
	Framework    string
	TemplatesFS  fs.FS
}

func (h *NewJavaHandler) SetTemplatesFS(fs fs.FS) {
	h.TemplatesFS = fs
}

const placeHolder = "PROJECT_NAME"

// Supported combinations of dependencies
var allowedCombinations = map[string][]string{
	"maven": {"", "springboot", "quarkus"}, // Maven allows no framework, Spring Boot, or Quarkus
}

// GetAllowedCombinations exposes the allowed build tool and framework combinations.
func GetAllowedCombinations() map[string][]string {
	return allowedCombinations
}

func (handler *NewJavaHandler) evaluateDependencies() error {
	// Default values
	handler.BuildTool = "maven"
	handler.Framework = "" // Default: No specific framework

	for _, dependency := range handler.Dependencies {
		lowerDep := strings.ToLower(dependency)
		if !isValidDependency(lowerDep) {
			return fmt.Errorf("unsupported dependency '%s'. Allowed dependencies are: %s",
				dependency, strings.Join(getAllowedDependencies(), ", "))
		}

		switch lowerDep {
		case "spring", "springboot":
			handler.Framework = "springboot"
		case "quarkus":
			handler.Framework = "quarkus"
		case "gradle":
			handler.BuildTool = "gradle"
		case "maven":
			handler.BuildTool = "maven"
		}
	}

	// Validate the combination of build tool and framework
	if err := validateCombination(handler.BuildTool, handler.Framework); err != nil {
		return err
	}

	return nil
}

func validateCombination(buildTool, framework string) error {
	validFrameworks, ok := allowedCombinations[buildTool]
	if !ok {
		return fmt.Errorf("unsupported build tool '%s'. Allowed build tools are: %s",
			buildTool, strings.Join(getAllowedBuildTools(), ", "))
	}

	// Check if the framework is allowed with the selected build tool
	for _, validFramework := range validFrameworks {
		if framework == validFramework {
			return nil
		}
	}

	return fmt.Errorf("unsupported combination: build tool '%s' does not support framework '%s'. Allowed frameworks for '%s' are: %s",
		buildTool, framework, buildTool, strings.Join(validFrameworks, ", "))
}

// Helper function to get all allowed dependencies
func getAllowedDependencies() []string {
	dependencies := make(map[string]struct{})
	for tool, frameworks := range allowedCombinations {
		dependencies[tool] = struct{}{}
		for _, framework := range frameworks {
			if framework != "" {
				dependencies[framework] = struct{}{}
			}
		}
	}

	// Convert map to a sorted slice
	result := make([]string, 0, len(dependencies))
	for dep := range dependencies {
		result = append(result, dep)
	}
	return result
}

// Helper function to get all allowed build tools
func getAllowedBuildTools() []string {
	buildTools := make([]string, 0, len(allowedCombinations))
	for tool := range allowedCombinations {
		buildTools = append(buildTools, tool)
	}
	return buildTools
}

func isValidDependency(dependency string) bool {
	for _, allowed := range getAllowedDependencies() {
		if allowed == dependency {
			return true
		}
	}
	return false
}

func (handler *NewJavaHandler) Run(projectName string) error {
	// Evaluate dependencies to determine BuildTool and Framework
	if err := handler.evaluateDependencies(); err != nil {
		return err
	}

	if handler.BuildTool == "" {
		return fmt.Errorf("invalid configuration: Build Tool not specified")
	}

	projectDirOnHost, err := utils.PrepareProjectDir(projectName)
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
	case "":
		// Default case: No specific framework
		return handler.setupDefaultMavenProject(projectDirOnHost, projectName)
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
	case "":
		// Default case: No specific framework
		return fmt.Errorf("setting up a plain Java project for Gradle is not yet implemented")
	case "quarkus":
		return fmt.Errorf("setting up a Quarkus project with Gradle is not yet implemented")
	default:
		return fmt.Errorf("unsupported framework '%s' for Gradle", handler.Framework)
	}
}

func (handler *NewJavaHandler) setupDefaultMavenProject(projectDirOnHost, projectName string) error {
	filesThatNeedProjectNameAdjustedOnce := []string{"docker-compose.dev.yml", "Makefile", "README.md"}
	filesThatNeedProjectNameAdjustedEverywhere := []string{"README.md"}
	filesThatNeedToBeRemoved := []string{"build.Dockerfile", "create_java_project.sh"}

	templatesPath := filepath.Join("templates", handler.Language, handler.BuildTool, "default")
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

	return fmt.Errorf("setting up a plain Java project for Maven is not implemented fully... a docker image named 'maven-project-generator:latest' is on your host and isn't cleaned up (can be done by running: 'docker image rm maven-project-generator:latest')")
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
