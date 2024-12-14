package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"craft/internal/constants"
)

var (
	operation string
	language  string
)

// Execute initializes and runs the root command.
func Execute() error {
	rootCmd := NewRootCmd()
	return rootCmd.Execute()
}

type Handler interface {
	Run() error
}

// NewRootCmd creates the root command.
func NewRootCmd() *cobra.Command {
	programName := strings.TrimPrefix(os.Args[0], "./")

	rootCmd := &cobra.Command{
		Use:   programName,
		Short: "A CLI tool with autocompletion support",
		Long:  fmt.Sprintf("This tool provides operations and language management with autocompletion for different shells. Run `%s help` for details.", programName),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Check if both flags are provided
			if operation != "" && language != "" {

				if err := validateFlags(operation, language); err != nil {
					return err
				}

				fmt.Printf("Performing '%s' operation in '%s' language\n", operation, language)
				handler, err := getHandler(operation, language)
				if err != nil {
					return err
				}
				handler.Run()

				return nil
			}

			return cmd.Help()
		},
		SilenceUsage: true,
	}

	rootCmd.Flags().StringVarP(&operation, "operation", "o", "", fmt.Sprintf("Specify the operation to perform (Allowed: %s)", strings.Join(constants.AllowedOperations, ", ")))
	rootCmd.Flags().StringVarP(&language, "language", "l", "", fmt.Sprintf("Specify the language for the project (Allowed: %s)", strings.Join(constants.AllowedLanguages, ", ")))

	return rootCmd
}

func getHandler(operation, language string) (Handler, error) {
	switch operation {
	case "new":
		{
			switch language {
			case "Go":
				return createGoHandler{}, nil
			}
		}

	}
	return nil, nil
}

type createGoHandler struct{}

func (h createGoHandler) Run() error {
	fmt.Println("Building a Go project...")
	// Logic for building a Go project
	return nil
}

// validateFlags ensures that the provided flag values are within the allowed set.
func validateFlags(operation, language string) error {
	if !constants.IsValidOperation(operation) {
		return fmt.Errorf("invalid operation: %s. Allowed operations are: %s", operation, strings.Join(constants.AllowedOperations, ", "))
	}

	if !constants.IsValidLanguage(language) {
		return fmt.Errorf("invalid language: %s. Allowed languages are: %s", language, strings.Join(constants.AllowedLanguages, ", "))
	}
	return nil
}
