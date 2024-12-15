package cmd

import (
	"craft/internal/constants"
	"craft/registry"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// NewNewCmd creates the `new` subcommand
func NewNewCmd() *cobra.Command {
	var currentDirName bool
	var name string

	// Dynamically fetch allowed languages for the `new` command
	allowedLanguages := registry.GetAllowedLanguages("new")
	allowedLanguagesText := strings.Join(allowedLanguages, ", ")

	cmd := &cobra.Command{
		Use:   "new <language>",
		Short: "Create a new project",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("missing required argument: <language>.\nSupported languages are: %v\n",
					allowedLanguagesText)
			}
			if len(args) > 1 {
				return fmt.Errorf("unexpected additional arguments: %v\n\n%s", args[1:], cmd.UsageString())
			}
			return nil
		},

		RunE: func(cmd *cobra.Command, args []string) error {
			language := args[0]
			err := constants.ValidateOperationAndLanguage("new", language)
			if err != nil {
				return err
			}

			if name == "" && !currentDirName {
				return fmt.Errorf("Run 'craft new <language>' with either --name <name> or --current-dir-name (-c) to specify the project name")
			}

			projectName := getProjectName(currentDirName, name)
			handler, err := registry.GetNewHandler(language)
			if err != nil {
				return err
			}
			handler.Run(projectName)

			return nil
		},

		SilenceUsage: true,
	}

	cmd.Flags().BoolVarP(&currentDirName, "current-dir-name", "c", false, "Passes the current directory name as the name for the new project")
	cmd.Flags().StringVarP(&name, "name", "n", "", "Specify the name of the new project")

	return cmd
}

// getProjectName extracts the project name based on flags or the current directory
func getProjectName(currentDirName bool, name string) string {
	if name != "" {
		return name
	}

	if currentDirName {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error fetching current directory:", err)
			return "run-app" // default name
		}
		return filepath.Base(wd)
	}

	return ""
}
