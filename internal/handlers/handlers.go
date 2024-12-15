package handlers

import (
	"craft/internal/common"
	gohandler "craft/internal/handlers/go"
	"fmt"
)

func GetNewHandler(language string) (common.NewHandler, error) {
	switch language {
	case "Go":
		return gohandler.NewGoHandler{Language: "go"}, nil
	default:
		return nil, fmt.Errorf("no 'new' handler found for language '%s'", language)
	}
}

func GetScanHandler(language string) (common.ScanHandler, error) {
	switch language {
	case "Go":
		return gohandler.ScanGoHandler{}, nil
	default:
		return nil, fmt.Errorf("no 'scan' handler found for language '%s'", language)
	}
}
