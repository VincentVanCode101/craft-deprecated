package utils

import (
	"archive/tar"
	"compress/gzip"
	"embed"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// UnzipAndUntar will unzip and untar the templates archive from the embed.FS and extract the contents to the specified directory.
func UnzipAndUntar(templatesFS embed.FS, destDir string) error {
	// Open the embedded tar.gz file.
	file, err := templatesFS.Open("templates.tar.gz")
	if err != nil {
		return fmt.Errorf("failed to open templates.tar.gz from embed.FS: %w", err)
	}
	defer file.Close()

	// Unzip the file using gzip.NewReader
	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer gzipReader.Close()

	// Untar the content using tar.NewReader
	tarReader := tar.NewReader(gzipReader)

	// Iterate over the tar contents and extract each file
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			return fmt.Errorf("failed to read tar entry: %w", err)
		}

		// Create the file path where the file should be extracted
		extractedFilePath := filepath.Join(destDir, header.Name)

		// If the entry is a directory, create the directory
		if header.Typeflag == tar.TypeDir {
			err := os.MkdirAll(extractedFilePath, os.ModePerm)
			if err != nil {
				return fmt.Errorf("failed to create directory %s: %w", extractedFilePath, err)
			}
			continue
		}

		// Otherwise, it's a file, so extract it
		outputFile, err := os.Create(extractedFilePath)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", extractedFilePath, err)
		}
		defer outputFile.Close()

		// Copy the contents of the file from the tar archive to the output file
		_, err = io.Copy(outputFile, tarReader)
		if err != nil {
			return fmt.Errorf("failed to extract file %s: %w", extractedFilePath, err)
		}
	}

	return nil
}
