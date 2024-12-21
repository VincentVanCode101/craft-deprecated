package cmd

import (
	"craft/internal/constants"
	"craft/internal/handlers"
	"craft/registry"
	"embed"
	"fmt"
	"strings"

	javahandler "craft/internal/handlers/java"

	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// NewNewCmd creates a new "new" command to generate a project scaffold for a specified language.
// It supports specifying a project name or using the current directory as the project name.
// Templates for the project are embedded in the provided templatesFS parameter.
func NewNewCmd(templatesFS embed.FS) *cobra.Command {
	var specifiedProjectName string
	var dependencies string

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
			showDependencies, _ := cmd.Flags().GetBool("show-dependencies")

			if showDependencies {
				if len(args) < 1 {
					return fmt.Errorf("please specify a language to show its dependencies")
				}

				titleCaser := cases.Title(language.English) // Proper Unicode casing
				language := args[0]
				dependenciesInfo := fetchSupportedDependenciesInfo(language, titleCaser)
				fmt.Println(dependenciesInfo)
				return nil
			}
			language := args[0]

			err := registry.ValidateOperationAndLanguage("new", language)
			if err != nil {
				return err
			}

			projectName := getProjectDetails(specifiedProjectName, language)

			rawDeps := strings.Split(dependencies, ",")
			deps := make([]string, 0, len(rawDeps))

			for _, dep := range rawDeps {
				trimmed := strings.TrimSpace(dep)
				if trimmed != "" {
					deps = append(deps, trimmed)
				}
			}

			handler, err := handlers.GetNewHandler(strings.ToLower(language), deps)
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

	cmd.Flags().StringVarP(&dependencies, "dependencies", "d", "", "Specify the dependencies or project type (e.g., -d maven-spring)")
	cmd.Flags().StringVarP(&specifiedProjectName, "name", "n", "", "Specify the project name (e.g. -n my-test-project)")
	cmd.Flags().Bool("show-dependencies", false, "Show supported dependencies for the specified language")

	return cmd
}

func getProjectDetails(specifiedProjectName, language string) string {
	if specifiedProjectName != "" {
		return specifiedProjectName
	}

	return fmt.Sprintf("%v-%v", constants.ToolName, language)
}

func fetchSupportedDependenciesInfo(language string, titleCaser cases.Caser) string {
	language = strings.ToLower(language)
	switch language {
	case "java":
		// Java-specific logic
		combinations := javahandler.GetAllowedCombinations()
		var sb strings.Builder
		sb.WriteString("Maven is the default build-tool for the java projects:\n\n")

		sb.WriteString("Supported Dependencies:\n")
		for buildTool, frameworks := range combinations {
			sb.WriteString(fmt.Sprintf("  for the build tool '%s' are:\n", buildTool))
			if len(frameworks) == 0 || (len(frameworks) == 1 && frameworks[0] == "") {
				sb.WriteString("    - No specific frameworks required\n")
			} else {
				for _, framework := range frameworks {
					if framework == "" {
						// sb.WriteString(fmt.Sprintf("    - No specific framework (just use %v without specifying anything)\n", titleCaser.String(buildTool)))
					} else {
						sb.WriteString(fmt.Sprintf("    - %s\n", titleCaser.String(framework)))
					}
				}
			}
		}
		return sb.String()
	case "go":
		// Go-specific logic (no dependencies)
		return "Supported Dependencies:\n  - None supported\n"
	default:
		return "No supported dependencies for this language."
	}
}
