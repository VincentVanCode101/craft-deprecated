package common

import (
	"craft/internal/utils"
	"fmt"
	"path/filepath"
)

// AdjustProjectNames replaces placeholders in the specified files with the project name.
// - projectHostDir: The base directory of the project.
// - onceFiles: Files where the placeholder should be replaced only once.
// - everywhereFiles: Files where the placeholder should be replaced throughout.
// - placeholder: The placeholder to be replaced.
// - projectName: The new project name to replace the placeholder with.
func AdjustProjectNames(projectHostDir string, onceFiles, everywhereFiles []string, placeholder, projectName string) error {
	// Replace placeholders in files where replacement happens only once
	for _, filePath := range onceFiles {
		if err := utils.ChangeWordInFile(filepath.Join(projectHostDir, filePath), placeholder, projectName, false); err != nil {
			return fmt.Errorf("error adjusting project name in file '%s': %v", filePath, err)
		}
	}

	// Replace placeholders in files where replacement happens everywhere
	for _, filePath := range everywhereFiles {
		if err := utils.ChangeWordInFile(filepath.Join(projectHostDir, filePath), placeholder, projectName, true); err != nil {
			return fmt.Errorf("error adjusting project name in file '%s': %v", filePath, err)
		}
	}
	return nil
}

func CleanupFiles(projectHostDir string, files []string) error {
	for _, file := range files {
		filePath := filepath.Join(projectHostDir, file)
		if err := utils.RemoveFileFromHost(filePath); err != nil {
			return fmt.Errorf("error removing file '%s': %v", filePath, err)
		}
	}
	return nil
}
