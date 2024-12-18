package handlers

import (
	"craft/internal/common"
	gohandler "craft/internal/handlers/go"
	javahandler "craft/internal/handlers/java"
	"fmt"
)

func GetNewHandler(language []string) (common.NewHandler, error) {
	if len(language) < 1 {
		return nil, fmt.Errorf("language must have at least one element")
	}

	switch language[0] {
	case "go":
		return &gohandler.NewGoHandler{Language: "go"}, nil

	case "java":
		if len(language) < 3 {
			return nil, fmt.Errorf("java language options must specify at least build tool and framework")
		}
		switch language[1] {
		case "maven":
			switch language[2] {
			case "noframework":
				return &javahandler.NewJavaHandler{Language: "java-maven-noframework"}, nil
			case "quarkus":
				return nil, fmt.Errorf("no handler for java maven quarkus implemented")
			case "springboot":
				return nil, fmt.Errorf("no handler for java maven springboot implemented")
			default:
				return nil, fmt.Errorf("no handler for java with build tool maven and framework '%s' implemented", language[2])
			}
		case "gradle":
			switch language[2] {
			case "quarkus":
				return nil, fmt.Errorf("no handler for java gradle quarkus implemented")
			case "springboot":
				return nil, fmt.Errorf("no handler for java gradle springboot implemented")
			case "noframework":
				return nil, fmt.Errorf("no handler for java gradle and no framework implemented")
			default:
				return nil, fmt.Errorf("no handler for java gradle framework '%s' implemented", language[2])
			}
		default:
			return nil, fmt.Errorf("no handler for java with build tool '%s' implemented", language[1])
		}

	default:
		return nil, fmt.Errorf("no 'new' handler found for language '%s'", language[0])
	}
}
