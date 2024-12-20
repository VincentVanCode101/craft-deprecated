package registry

var (
	AllowedOperationsWithLanguages = map[string][]string{
		"new": {"java", "go"},
	}
)

// GetAllowedLanguages returns the supported languages for a specific operation.
func GetAllowedLanguages(operation string) []string {
	if languages, exists := AllowedOperationsWithLanguages[operation]; exists {
		return languages
	}
	return []string{}
}
