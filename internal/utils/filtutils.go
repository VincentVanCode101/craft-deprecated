package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// CopyFile copies a file from src to dst
func CopyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer destFile.Close()

	bytesCopied, err := io.Copy(destFile, sourceFile)
	if err != nil {
		return fmt.Errorf("failed to copy file contents (bytes copied: %d): %w", bytesCopied, err)
	}

	err = destFile.Sync()
	if err != nil {
		return fmt.Errorf("failed to flush data to destination file: %w", err)
	}

	return nil
}

// ChangeProjectNameInFile replaces a placeholder with the given projectName in a file
func ChangeProjectNameInFile(fileName, projectName string) error {
	err := os.Chmod(fileName, 0644) // Grant read/write for owner, read for group/others
	if err != nil {
		return fmt.Errorf("failed to change file permissions for %s: %w", fileName, err)
	}

	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "PROJECT_NAME") {
			line = strings.Replace(line, "PROJECT_NAME", projectName, -1)
		}
		lines = append(lines, line)
	}

	err = scanner.Err()
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	// Open the file for writing (truncate mode)
	outputFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("error opening file for writing: %w", err)
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("error writing to file: %w", err)
		}
	}

	if err := writer.Flush(); err != nil {
		return fmt.Errorf("error flushing writer: %w", err)
	}

	return nil
}

// ListDirectoryContents lists all files and directories in the given path
func ListDirectoryContents(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("error reading directory: %w", err)
	}

	var results []string
	for _, entry := range entries {
		if entry.IsDir() {
			results = append(results, "[DIR] "+entry.Name())
		} else {
			results = append(results, "[FILE] "+entry.Name())
		}
	}
	return results, nil
}

// GetFilePath constructs the full file path given a base path and a file name
func GetFilePath(basePath, fileName string) string {
	return filepath.Join(basePath, fileName)
}

// GetFilePaths constructs a list of full file paths for multiple files
func GetFilePaths(basePath string, files []string) []string {
	var filePaths []string
	for _, file := range files {
		filePaths = append(filePaths, filepath.Join(basePath, file))
	}
	return filePaths
}
