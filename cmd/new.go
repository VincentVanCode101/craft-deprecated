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
	var useCurrentDirName bool
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

			if name == "" && !useCurrentDirName {
				return fmt.Errorf("Run 'craft new <language>' with either --name <name> or --current-dir-name (-c) to specify the project name.")
			}

			createDirectoryFor, projectName := getProjectDetails(useCurrentDirName, name)
			handler, err := registry.GetNewHandler(language)
			if err != nil {
				return err
			}
			handler.Run(createDirectoryFor, projectName)

			return nil
		},

		SilenceUsage: true,
	}

	cmd.Flags().BoolVarP(&useCurrentDirName, "current-dir-name", "c", false, "Use the current directory name for the new project. The new files will be created in the current directory without creating a new one.")
	cmd.Flags().StringVarP(&name, "name", "n", "", "Specify a name for the new project. A new directory with this name will be created, and the files will be placed inside it.")

	return cmd
}

func getProjectDetails(useCurrentDirName bool, name string) (bool, string) {
	if name != "" {
		return true, name
	}

	if useCurrentDirName {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error fetching current directory:", err)
			return true, "run-app" // default name
		}
		fmt.Println(filepath.Base(wd))
		return false, filepath.Base(wd)
	}

	return false, ""
}
