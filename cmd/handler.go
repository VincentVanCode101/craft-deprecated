package cmd

import "fmt"

// Handler interface defines the Run method for operations.
type Handler interface {
	Run() error
}

// getHandler returns a handler based on the operation and language.
func getHandler(operation, language string) (Handler, error) {
	switch operation {
	case "new":
		switch language {
		case "Go":
			return createGoHandler{}, nil
		}
	}
	return nil, fmt.Errorf("no handler found for operation '%s' and language '%s'", operation, language)
}

type createGoHandler struct{}

func (h createGoHandler) Run() error {
	fmt.Println("Building a Go project...")
	// Logic for building a Go project
	return nil
}
