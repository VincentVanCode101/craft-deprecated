package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// NewCompletionCmd creates the "completion" command.
func NewCompletionCmd() *cobra.Command {
	completionCmd := &cobra.Command{
		Use:   "completion",
		Short: "Generate completion script for your shell",
		Long:  `Generate shell completion scripts for your preferred shell environment.`,
		Run:   runCompletion,
	}

	completionCmd.Flags().String("shell", "", "Specify the shell type (bash, zsh, fish)")
	return completionCmd
}

func runCompletion(cmd *cobra.Command, args []string) {
	shell, _ := cmd.Flags().GetString("shell")
	programName := strings.TrimPrefix(os.Args[0], "./")

	if shell == "" {
		fmt.Println("Error: You must specify a shell using the --shell flag (e.g., bash, zsh, fish).")
		fmt.Println(cmd.UsageString())
		return
	}

	var path string
	switch shell {
	case "bash":
		path = "/etc/bash_completion.d/" + programName
	case "zsh":
		path = os.ExpandEnv("~/.zsh/completions/") + programName
	case "fish":
		path = os.ExpandEnv("~/.config/fish/completions/") + programName + ".fish"
	default:
		fmt.Println("Unsupported shell. Please specify --shell as 'bash', 'zsh', or 'fish'.")
		fmt.Println(cmd.UsageString())
		return
	}

	file, err := os.Create(path)
	if err != nil {
		fmt.Printf("Error saving the script to %s: %v\n", path, err)
		return
	}
	defer file.Close()

	switch shell {
	case "bash":
		err = cmd.Root().GenBashCompletion(file)
	case "zsh":
		err = cmd.Root().GenZshCompletion(file)
	case "fish":
		err = cmd.Root().GenFishCompletion(file, true)
	}

	if err != nil {
		fmt.Printf("Error generating %s completion script: %v\n", shell, err)
		return
	}

	fmt.Printf("Completion script successfully saved to %s\n", path)
}
