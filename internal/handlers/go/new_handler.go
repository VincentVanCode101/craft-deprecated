package gohandler

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"

	"craft/internal/common"
	"craft/internal/utils"
)

type NewGoHandler struct {
	Dependencies        []string
	Language            string
	TemplatesFileSystem fs.FS
}

func (h *NewGoHandler) SetTemplatesFS(fileSystem fs.FS) {
	h.TemplatesFileSystem = fileSystem
}

const (
	templateFileSuffix     = ".template"
	dotFileNotationPrefix  = "DOT"
	dotFilePrefix          = "."
	projectNamePlaceholder = "{PROJECT_NAME}"
)

var (
	filesThatNeedProjectNameAdjustedOnce       = []string{"go.mod.template", "Makefile"}
	filesThatNeedProjectNameAdjustedEverywhere = []string{"README.md", "docker-compose.dev.yml"}
)

func (h *NewGoHandler) Run(projectName string) error {
	var projectHostDir string
	var err error

	projectHostDir, err = utils.PrepareProjectDir(projectName)
	if err != nil {
		fmt.Printf("Failed to get current working directory: %v\n", err)
		return err
	}
	fmt.Printf("The project directory: %v\n\n", projectHostDir)

	languageTemplatePath := filepath.Join("templates", h.Language)

	err = utils.CopyDirFromFS(h.TemplatesFileSystem, languageTemplatePath, projectHostDir)
	if err != nil {
		fmt.Printf("Error copying files from the embedded folder: %v to host: %v -> error: %v\n", languageTemplatePath, projectHostDir, err)
		return err
	}
	if err := h.adjustProjectNames(projectHostDir, filesThatNeedProjectNameAdjustedOnce, filesThatNeedProjectNameAdjustedEverywhere, projectName); err != nil {
		return err
	}

	if err := h.adjustProjectNames(projectHostDir, filesThatNeedProjectNameAdjustedOnce, filesThatNeedProjectNameAdjustedEverywhere, projectName); err != nil {
		return err
	}

	dotFileCandidates, err := utils.ListFilesWithPattern(h.TemplatesFileSystem, languageTemplatePath, dotFileNotationPrefix)
	if err != nil {
		return err
	}

	for _, filePath := range dotFileCandidates {
		hostFilePath := path.Join(projectHostDir, filePath)
		renamedFilePath := strings.ReplaceAll(hostFilePath, dotFileNotationPrefix, dotFilePrefix)

		err := os.Rename(hostFilePath, renamedFilePath)

		if err != nil {
			fmt.Printf("Error renaming file %v to remove %v: %v\n", hostFilePath, dotFileNotationPrefix, err)
			return err
		}
	}

	templateFiles, err := utils.ListFilesWithPattern(h.TemplatesFileSystem, languageTemplatePath, templateFileSuffix)
	if err != nil {
		return err
	}
	for _, filePath := range templateFiles {
		hostFilePath := path.Join(projectHostDir, filePath)
		cleanedFilePath := strings.TrimSuffix(hostFilePath, templateFileSuffix)

		err := os.Rename(hostFilePath, cleanedFilePath)
		if err != nil {
			fmt.Printf("Error removing suffix %v in %v: %v\n", templateFileSuffix, hostFilePath, err)
			return err
		}
	}

	return nil
}

func (handler *NewGoHandler) adjustProjectNames(projectHostDir string, onceFiles, everywhereFiles []string, projectName string) error {
	return common.AdjustProjectNames(projectHostDir, onceFiles, everywhereFiles, projectNamePlaceholder, projectName)
}
