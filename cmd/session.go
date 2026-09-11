package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var sessionCmd = &cobra.Command{
	Use:   "session",
	Short: "Start a new magicTerminal session",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use 'magic session --help' to see available options.")
	},
}

func init() {
	// This links session directly to the root command automatically
	rootCmd.AddCommand(sessionCmd)
}