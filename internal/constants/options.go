package constants

import (
	"craft/internal/utils"
	"craft/registry"
	"fmt"
	"strings"
)

func ValidateOperationAndLanguage(operation, language string) error {
	lowerCaseLanguage := strings.ToLower(language)
	allowedLanguages := registry.GetAllowedLanguages(operation)

	if len(allowedLanguages) == 0 {
		return fmt.Errorf("operation '%s' is not allowed. Allowed operations are: %v", operation, utils.Keys(registry.AllowedOperationsWithLanguages))
	}

	if !utils.ContainsStringInsensitive(allowedLanguages, lowerCaseLanguage) {
		return fmt.Errorf(
			"operation '%s' cannot be performed with language '%s'.\nAllowed languages for this operation are: {%v}",
			operation, language, strings.Join(allowedLanguages, ", "))
	}

	return nil
}
