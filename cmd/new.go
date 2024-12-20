package cmd

import (
	"craft/internal/constants"
	"craft/internal/handlers"
	"craft/registry"
	"embed"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// NewNewCmd creates a new "new" command to generate a project scaffold for a specified language.
// It supports specifying a project name or using the current directory as the project name.
// Templates for the project are embedded in the provided templatesFS parameter.
func NewNewCmd(templatesFS embed.FS) *cobra.Command {
	var specifiedProjectName string

	allowedLanguages := registry.GetAllowedLanguages("new")
	allowedLanguagesText := strings.Join(allowedLanguages, ", ")

	cmd := &cobra.Command{
		Use:   "new <language>",
		Short: "Create a new project",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("missing required argument: <language>.\nSupported languages are: %v",
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

			projectName := getProjectDetails(specifiedProjectName, language)

			languageStrings := strings.Split(strings.ToLower(language), "-")
			handler, err := handlers.GetNewHandler(languageStrings)
			if err != nil {
				return err
			}
			handler.SetTemplatesFS(&templatesFS)
			err = handler.Run(projectName)
			if err != nil {
				return err
			}

			return nil
		},

		SilenceUsage: true,
	}

	cmd.Flags().StringVarP(&specifiedProjectName, "name", "n", "", "Specify a name for the new project. A new directory with this name will be created, and the files will be placed inside it.")

	return cmd
}

func getProjectDetails(specifiedProjectName, language string) string {
	if specifiedProjectName != "" {
		return specifiedProjectName
	}

	return fmt.Sprintf("%v-%v", constants.Tool_name, language)
}
