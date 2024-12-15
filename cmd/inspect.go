package cmd

import (
	"craft/internal/constants"
	"fmt"

	"github.com/spf13/cobra"
)

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
	for operation, languages := range constants.AllowedOperationsWithLanguages {
		fmt.Printf("- %s: %v\n", operation, languages)
	}
}
