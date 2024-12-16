package common

import "io/fs"

// NewHandler defines the interface for creating a new project.
// It includes methods for running the handler and setting a template filesystem.
type NewHandler interface {
	Run(createDirectoryFor bool, projectName string) error
	SetTemplatesFS(fs fs.FS)
}
