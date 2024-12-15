package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "craft",
		Short: "A CLI tool with autocompletion support",
		Long:  "This tool provides operations and language management with autocompletion for different shells. Run `craft help` for details.",
	}

	rootCmd.AddCommand(NewNewCmd())
	rootCmd.AddCommand(NewInspectCmd())

	return rootCmd
}
