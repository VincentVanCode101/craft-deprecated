package common

import "io/fs"

type NewHandler interface {
	Run(createDirectoryFor bool, projectName string) error
	SetTemplatesFS(fs fs.FS)
}

type ScanHandler interface {
	Run() error
}
