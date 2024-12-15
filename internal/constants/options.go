package constants

import (
	"craft/internal"
	"fmt"
)

var (
	AllowedOperationsWithLanguages = map[string][]string{
		"new":  {"Go", "Java"},
		"scan": {"Go"},
	}
)

// ValidateOperationAndLanguage checks if the operation and language are valid
// and returns detailed error messages.
func ValidateOperationAndLanguage(operation, language string) error {
	arr, foundOperation := AllowedOperationsWithLanguages[operation]
	if !foundOperation {
		return fmt.Errorf("operation '%s' is not allowed. Allowed operations are: %v", operation, internal.Keys(AllowedOperationsWithLanguages))
	}

	if !internal.Contains(arr, language) {
		return fmt.Errorf("operation '%s' cannot be performed with language '%s'. Allowed languages for this operation are: %v", operation, language, arr)
	}

	return nil
}
