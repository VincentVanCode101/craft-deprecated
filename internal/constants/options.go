package constants

var (
	AllowedOperations = []string{"new"}
	AllowedLanguages  = []string{"Go", "Rust"}
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

// IsValidOperation checks if the operation is allowed.
func IsValidOperation(operation string) bool {
	return contains(AllowedOperations, operation)
}

// IsValidLanguage checks if the language is allowed.
func IsValidLanguage(language string) bool {
	return contains(AllowedLanguages, language)
}
