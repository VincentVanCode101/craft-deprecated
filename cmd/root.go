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
		RunE: func(cmd *cobra.Command, args []string) error {
			if inspect {
				showAllowedOptions()
				return nil
			}

			argLen := len(args)
			switch {
			case argLen == 0:
				return cmd.Help()
			case argLen == 1:
				return fmt.Errorf("craft requires at least 2 arguments: <operation> <language>")
			case argLen == 2:
				operation := args[0]
				language := args[1]

				err := constants.ValidateOperationAndLanguage(operation, language)
				if err != nil {
					return err
				}

				if operation == "new" {
					if name == "" && !currentDirName {
						return fmt.Errorf("Run 'craft new <language>' with either --name <name> or --current-dir-name (-c) to specify the project name")
					}
					projectName := getProjectName(currentDirName, name)
					fmt.Printf("Performing '%s' operation for language '%s' with project name '%s'\n", operation, language, projectName)
				} else {
					fmt.Printf("Performing '%s' operation for language '%s'\n", operation, language)
				}

				return nil
			case argLen > 2:
				return fmt.Errorf("craft does not take more than 2 arguments. You provided %d: %v", argLen, args)
			}
			return nil
		},

		SilenceUsage: true,
	}

	rootCmd.PersistentFlags().BoolVarP(&inspect, "inspect", "i", false, "Show allowed operations and languages")
	rootCmd.Flags().BoolVarP(&currentDirName, "current-dir-name", "c", false, "Passes the current directory name as the name for the new project (only for 'new' operation)")
	rootCmd.Flags().StringVarP(&name, "name", "n", "", "Specify the name of the new project (only for 'new' operation)")

	return rootCmd
}

func getProjectName(currentDirName bool, name string) string {
	if name != "" {
		return name
	}

	if currentDirName {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error fetching current directory:", err)
			return "run-app" // default name (pondering: is this a good default project name?)
		}
		return filepath.Base(wd)
	}

	return ""
}

func showAllowedOptions() {
	fmt.Println("Allowed Operations and Languages:")
	for operation, languages := range constants.AllowedOperationsWithLanguages {
		fmt.Printf("- %s: %v\n", operation, languages)
	}
}
