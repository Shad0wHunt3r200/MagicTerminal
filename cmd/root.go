package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "magic",
	Short: "MagicTerminal root command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("👋 Welcome to MagicTerminal.\nType 'magic --help' to see available options.")
	},
}

// Execute is called by main.go to start the CLI engine
func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}