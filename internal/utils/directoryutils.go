package utils

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// PrepareProjectDir prepares a project directory.
// If `createDir` is true, it creates a new directory for the project.
// If `createDir` is false, it returns the current working directory.
// `projectName` is used to name the new directory if created.
func PrepareProjectDir(createDir bool, projectName string) (string, error) {
	currentPwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("could not get current directory: %w", err)
	}

	if createDir {
		projectDir := filepath.Join(currentPwd, projectName)

		err := os.Mkdir(projectDir, 0755)
		if err != nil {
			return "", fmt.Errorf("could not create project directory: %w", err)
		}
		return projectDir, nil
	}

	return currentPwd, nil
}

func CopyDirFromFS(fsys fs.FS, sourceDir, destDir string) error {
	err := fs.WalkDir(fsys, sourceDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("error walking directory: %w", err)
			return err
		}

		realPath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			fmt.Errorf("error calculating relative path: %w", err)
			return err
		}
		targetPath := filepath.Join(destDir, realPath)

		if d.IsDir() {
			err = os.MkdirAll(targetPath, directoryPermissions)
			if err != nil {
				fmt.Errorf("error creating directory %q: %w", targetPath, err)
				return err
			}
		} else {
			err = CopyFileFromFS(fsys, path, targetPath)
			if err != nil {
				fmt.Errorf("error copying file: %w", err)
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// CopyDirIntoDir copies a source directory into a destination directory
func CopyDirIntoDir(sourceDir, destinationDir string) error {
	sourceEntries, err := os.ReadDir(sourceDir)
	if err != nil {
		return fmt.Errorf("error reading directory %s: %w", sourceDir, err)
	}

	destinationPath := filepath.Join(destinationDir, filepath.Base(sourceDir))
	if err := os.MkdirAll(destinationPath, os.ModePerm); err != nil {
		return fmt.Errorf("error creating directory %s: %w", destinationPath, err)
	}

	for _, entry := range sourceEntries {
		sourcePath := filepath.Join(sourceDir, entry.Name())
		destPath := filepath.Join(destinationPath, entry.Name())

		if entry.IsDir() {
			if err := CopyDirIntoDir(sourcePath, destinationPath); err != nil {
				return err
			}
		} else {
			if err := CopyFile(sourcePath, destPath); err != nil {
				return err
			}
		}
	}

	return nil
}

// GetAllDirs retrieves all directory paths from a given directory on the host filesystem
func GetAllDirs(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("error reading directory %s: %w", dir, err)
	}

	var dirs []string
	for _, entry := range entries {
		if entry.IsDir() {
			absPath, err := filepath.Abs(filepath.Join(dir, entry.Name()))
			if err != nil {
				return nil, fmt.Errorf("error getting absolute path for %s: %w", entry.Name(), err)
			}
			dirs = append(dirs, absPath)
		}
	}

	return dirs, nil
}

func CopyAllOnePathUpAndRemoveDir(dirPath string) error {
	parentPath := filepath.Dir(dirPath)

	err := CopyAllEntries(dirPath, parentPath)
	if err != nil {
		return err
	}

	err = os.RemoveAll(dirPath)
	if err != nil {
		return err
	}

	return nil
}
