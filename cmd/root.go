package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "craft",
		Short: "A CLI tool to help bootstrap new Projects ",
		Long:  "This tool helps create new projects quickly by generating boilerplate code for a specified language or framework. Everything is configured to ensure the project runs seamlessly in a Docker container. Run craft help for more details.",
	}

	rootCmd.AddCommand(NewNewCmd())
	rootCmd.AddCommand(NewInspectCmd())

	return rootCmd
}
