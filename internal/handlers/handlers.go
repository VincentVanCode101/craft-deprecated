package handlers

import (
	"craft/internal/common"
	gohandler "craft/internal/handlers/go"
	javahandler "craft/internal/handlers/java"
	"fmt"
)

func GetNewHandler(language string) (common.NewHandler, error) {
	switch language {
	case "Go":
		return &gohandler.NewGoHandler{Language: "go"}, nil
	case "Java":
		return &javahandler.NewJavaHandler{Language: "java"}, nil
	default:
		return nil, fmt.Errorf("no 'new' handler found for language '%s'", language)
	}
}
