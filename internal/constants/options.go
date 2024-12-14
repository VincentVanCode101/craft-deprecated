package constants

import "fmt"

var (
	AllowedOperationsWithLanguages = map[string][]string{
		"new":  {"Go", "Java"},
		"scan": {"Go"},
	}
)

// contains checks if a slice contains a specific string.
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// ValidateOperationAndLanguage checks if the operation and language are valid
// and returns detailed error messages.
func ValidateOperationAndLanguage(operation, language string) error {
	arr, foundOperation := AllowedOperationsWithLanguages[operation]
	if !foundOperation {
		return fmt.Errorf("operation '%s' is not allowed. Allowed operations are: %v", operation, allowedOperations())
	}

	if !contains(arr, language) {
		return fmt.Errorf("operation '%s' cannot be performed with language '%s'. Allowed languages for this operation are: %v", operation, language, arr)
	}

	return nil
}

func allowedOperations() []string {
	keys := make([]string, 0, len(AllowedOperationsWithLanguages))
	for key := range AllowedOperationsWithLanguages {
		keys = append(keys, key)
	}
	return keys
}
