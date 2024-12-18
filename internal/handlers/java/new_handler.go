package javahandler

import "io/fs"

type NewJavaHandler struct {
	Language    string
	TemplatesFS fs.FS
}

func (h *NewJavaHandler) SetTemplatesFS(fs fs.FS) {
	h.TemplatesFS = fs
}

const (
	placeHolder = "PROJECT_NAME"
)

var (
	filesThatNeedProjectNameAdjusted = []string{"go.mod.template", "docker-compose.dev.yml", "Makefile"}
)

func (handler *NewJavaHandler) Run(createDirectoryFor bool, projectName string) error {
	// Steps:
	// 1. Copy everything over
	// 2. execute ./create_java_project.sh
	// 3. if --name was passed, copy everything into the PROJECT_NAME folder
	// 4. if -c was passed, copy everything from PROJECT_NAME folder one layer up
	// & remove PROJECT_NAME folder
	// 5. change PROJECT_NAME in everyfile to projectName
	// 6. remove the ./create_java_project.sh && build.Dockerfile
	// 7. remove docker image (DOCKER_IMAGE_NAME="maven-project-generator") from host
	return fs.ErrClosed
}
