package cmd

import (
	"craft/internal/constants"
	"fmt"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	programName := "craft"

	var showOptions bool // Flag variable

	rootCmd := &cobra.Command{
		Use:   fmt.Sprintf("%s <operation> <language>", programName),
		Short: "A CLI tool with autocompletion support",
		Long:  fmt.Sprintf("This tool provides operations and language management with autocompletion for different shells. Run `%s help` for details.", programName),
		Args:  cobra.MaximumNArgs(2), // Allow 0 to 2 arguments
		RunE: func(cmd *cobra.Command, args []string) error {
			// Check if the showOptions flag is set
			if showOptions {
				showAllowedOptions()
				return nil
			}

			if len(args) == 0 {
				return cmd.Help()
			}

			// Ensure exactly 2 arguments for operation and language
			if len(args) != 2 {
				return fmt.Errorf("invalid usage: you must specify both an operation and a language. Run '%s --help' for details", programName)
			}

			operation := args[0]
			language := args[1]

			// Validate operation and language
			err := constants.ValidateOperationAndLanguage(operation, language)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Operation and language are valid!")
			}

			handler, err := getHandler(operation, language)
			if err != nil {
				return err
			}
			handler.Run()

			return nil
		},
		SilenceUsage: true,
	}

	rootCmd.Flags().BoolVarP(&showOptions, "show-options", "s", false, "Show allowed operations and languages")

	return rootCmd
}

func showAllowedOptions() {
	fmt.Println("Allowed Operations and Languages:")
	for operation, languages := range constants.AllowedOperationsWithLanguages {
		fmt.Printf("- %s: %v\n", operation, languages)
	}
}
