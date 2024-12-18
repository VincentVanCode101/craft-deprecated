package constants

import (
	"craft/internal/utils"
	"fmt"
	"strings"
)

var (
	AllowedOperationsWithLanguages = map[string][]string{
		"new": {"Go", "Java-Maven-NoFramework"},
	}
)

// ValidateOperationAndLanguage checks if the operation and language are valid
// and returns detailed error messages.
func ValidateOperationAndLanguage(operation, language string) error {
	lowerCaseLanguage := strings.ToLower(language) // Normalize input to lowercase

	arr, foundOperation := AllowedOperationsWithLanguages[operation]
	if !foundOperation {
		return fmt.Errorf("operation '%s' is not allowed. Allowed operations are: %v", operation, utils.Keys(AllowedOperationsWithLanguages))
	}

	lowerCaseLanguages := make([]string, len(arr))
	for i, lang := range arr {
		lowerCaseLanguages[i] = strings.ToLower(lang)
	}

	if !utils.Contains(lowerCaseLanguages, lowerCaseLanguage) {
		return fmt.Errorf("operation '%s' cannot be performed with language '%s'.\nAllowed languages for this operation are: {%v}", operation, language, strings.Join(arr, ", "))
	}

	return nil
}
