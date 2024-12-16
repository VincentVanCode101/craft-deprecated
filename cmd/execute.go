package cmd

import "embed"

// Execute initializes and runs the root command.
func Execute(templatesFS embed.FS) error {
	rootCmd := NewRootCmd(templatesFS)
	return rootCmd.Execute()
}
