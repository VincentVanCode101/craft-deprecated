package registry

import (
	"fmt"
)

// Handler interfaces
type NewHandler interface {
	Run(createDirectoryFor bool, projectName string) error
}

type ScanHandler interface {
	Run() error
}

// Registry entry struct
type OperationEntry struct {
	Languages []string
	Handler   interface{}
}

// Centralized registry for operations
var OperationsRegistry = map[string]OperationEntry{
	"new": {
		Languages: []string{"Go", "Java"},
		Handler:   GetNewHandler,
	},
	"scan": {
		Languages: []string{"Go"},
		Handler:   GetScanHandler,
	},
}

func GetAllowedLanguages(operation string) []string {
	if entry, exists := OperationsRegistry[operation]; exists {
		return entry.Languages
	}
	return []string{}
}

// GetNewHandler returns a NewHandler for the specified language
func GetNewHandler(language string) (NewHandler, error) {
	switch language {
	case "Go":
		return newGoHandler{}, nil
	default:
		return nil, fmt.Errorf("no 'new' handler found for language '%s'", language)
	}
}

// getScanHandler returns a ScanHandler for the specified language
func GetScanHandler(language string) (ScanHandler, error) {
	switch language {
	case "Go":
		return scanGoHandler{}, nil
	default:
		return nil, fmt.Errorf("no 'scan' handler found for language '%s'", language)
	}
}

// Handlers for "new"
type newGoHandler struct{}

func (h newGoHandler) Run(createDirectoryFor bool, projectName string) error {
	fmt.Printf("Creating a new Go project with name '%s'...\n", projectName)
	return nil
}

// Handlers for "scan"
type scanGoHandler struct{}

func (h scanGoHandler) Run() error {
	fmt.Println("Scanning a Go project...")
	return nil
}
