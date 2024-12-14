package cmd

import (
	"craft/internal/constants"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	var inspect bool
	var currentDirName bool
	var name string

	rootCmd := &cobra.Command{
		Use:   "craft <operation> <language>",
		Short: "A CLI tool with autocompletion support",
		Long:  "This tool provides operations and language management with autocompletion for different shells. Run `craft help` for details.",
		Args:  cobra.MinimumNArgs(2), // Require at least two arguments: operation and language
		RunE: func(cmd *cobra.Command, args []string) error {
			if inspect {
				showAllowedOptions()
				// Exit after showing allowed options
				return nil
			}

			operation := args[0]
			language := args[1]

			// Validate the operation and language
			err := constants.ValidateOperationAndLanguage(operation, language)
			if err != nil {
				return err
			}

			if operation == "new" {
				// For "new" operation, validate project name or directory
				if name == "" && !currentDirName {
					return fmt.Errorf("Run 'craft new <language>' with either --name <name> or --current-dir-name (-c) to specify the project name")
				}
				projectName := getProjectName(currentDirName, name)
				fmt.Printf("Performing '%s' operation for language '%s' with project name '%s'\n", operation, language, projectName)
			} else {
				// For other operations, no additional flags are needed
				fmt.Printf("Performing '%s' operation for language '%s'\n", operation, language)
			}

			// Add the actual operation handling logic here
			return nil
		},
		SilenceUsage: true,
	}

	// Define global flags
	rootCmd.Flags().BoolVarP(&inspect, "inspect", "i", false, "Show allowed operations and languages")
	rootCmd.Flags().BoolVarP(&currentDirName, "current-dir-name", "c", false, "Passes the current directory name as the name for the new project (only for 'new' operation)")
	rootCmd.Flags().StringVarP(&name, "name", "n", "", "Specify the name of the new project (only for 'new' operation)")

	return rootCmd
}

func getProjectName(currentDirName bool, name string) string {
	if name != "" {
		// Use the name provided by the --name flag
		return name
	}

	if currentDirName {
		// Get the name of the current working directory
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error fetching current directory:", err)
			return "default-project-name"
		}
		return filepath.Base(wd)
	}

	// This case should never be reached due to the validation in RunE
	return ""
}

func showAllowedOptions() {
	fmt.Println("Allowed Operations and Languages:")
	for operation, languages := range constants.AllowedOperationsWithLanguages {
		fmt.Printf("- %s: %v\n", operation, languages)
	}
}
