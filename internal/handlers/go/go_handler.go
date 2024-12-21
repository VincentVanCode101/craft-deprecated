package gohandler

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"craft/internal/common"
	"craft/internal/constants"
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

	if err := h.copyTemplateFilesToHost(languageTemplatePath, projectHostDir); err != nil {
		return err
	}

	if err := h.adjustProjectNames(projectHostDir, filesThatNeedProjectNameAdjustedOnce, filesThatNeedProjectNameAdjustedEverywhere, projectName); err != nil {
		return err
	}

	dotFileCandidates, err := utils.ListFilesWithPattern(h.TemplatesFileSystem, languageTemplatePath, constants.DotFileNotationPrefix)
	if err != nil {
		return err
	}

	if err := utils.RenameFilesWithPrefix(dotFileCandidates, projectHostDir, constants.DotFileNotationPrefix, constants.DotFilePrefix); err != nil {
		fmt.Printf("Error renaming dot files: %v\n", err)
		return err
	}

	templateFiles, err := utils.ListFilesWithPattern(h.TemplatesFileSystem, languageTemplatePath, constants.TemplateFileSuffix)
	if err != nil {
		return err
	}

	if err := utils.TrimFileSuffix(templateFiles, projectHostDir, constants.TemplateFileSuffix); err != nil {
		fmt.Printf("Error cleaning template file suffixes: %v\n", err)
		return err
	}

	return nil
}

func (h *NewGoHandler) copyTemplateFilesToHost(languageTemplatePath, projectHostDir string) error {
	if err := utils.CopyDirFromFS(h.TemplatesFileSystem, languageTemplatePath, projectHostDir); err != nil {
		return fmt.Errorf("error copying files from template path: %v", err)
	}
	return nil
}

func (handler *NewGoHandler) adjustProjectNames(projectHostDir string, onceFiles, everywhereFiles []string, projectName string) error {
	return common.AdjustProjectNames(projectHostDir, onceFiles, everywhereFiles, constants.ProjectNamePlaceholder, projectName)
}
