package cmd

// Execute initializes and runs the root command.
func Execute() error {
	rootCmd := NewRootCmd()
	return rootCmd.Execute()
}
