package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func main() {
	// root command setup
	var rootCmd = &cobra.Command{
		Use:   "magic",
		Short: "MagicTerminal main command",
		Run: func(cmd *cobra.Command, args []string) {
			// This runs if user types 'magic' with no flags or subcommands
			fmt.Println("👋 Welcome to MagicTerminal.\nType 'magic --help' to see available options.")
		},
	}

	var dirCheckCmd = &cobra.Command{
		Use:   "dir-check",
		Short: "Check directory context",
		Run: func(cmd *cobra.Command, args []string) {
			gitCmd := exec.Command("git", "status", "--porcelain")
			_, err := gitCmd.Output()

			if err != nil {
				fmt.Println("📂 Context: Standard Directory")
				return
			}

			fmt.Println("🌲 Context: Git Repository")
		},
	}

	// Link the sub-command to the root command
	rootCmd.AddCommand(dirCheckCmd)

	// Start the Cobra engine
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
