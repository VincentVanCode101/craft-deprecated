package cmd

import (
	"craft/registry"
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// NewInspectCmd creates a new "inspect" command that displays allowed operations and languages.
func NewInspectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "inspect",
		Short: "Show allowed operations and languages",
		RunE: func(cmd *cobra.Command, args []string) error {
			showAllowedOptions()
			return nil
		},
	}
	return cmd
}

func showAllowedOptions() {
	fmt.Println("Allowed Operations and Languages:")

	titleCaser := cases.Title(language.English) // Use English for title casing

	for operation, languages := range registry.AllowedOperationsWithLanguages {
		fmt.Printf("- Operation: %s\n", titleCaser.String(operation))
		for _, language := range languages {
			fmt.Printf("  * Language: %s\n", titleCaser.String(language))
			fmt.Printf("    Run 'craft %s %s --help' to see the available dependencies.\n",
				operation, language)
		}
	}
}
