package cmd

import (
	"craft/internal/constants"
	"fmt"

	"github.com/spf13/cobra"
)

// NewInspectCmd creates the "inspect" command.
func NewInspectCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "inspect",
		Short: "Show allowed operations and languages",
		Run: func(cmd *cobra.Command, args []string) {
			showAllowedOptions()
		},
	}
}

// showAllowedOptions displays allowed operations and languages.
func showAllowedOptions() {
	fmt.Println("Allowed Operations and Languages")
	fmt.Printf("Operations: %v\n", constants.AllowedOperations)
	fmt.Printf("Languages: %v\n", constants.AllowedLanguages)
}
