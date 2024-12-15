package common

type NewHandler interface {
	Run(createDirectoryFor bool, projectName string) error
}

type ScanHandler interface {
	Run() error
}
