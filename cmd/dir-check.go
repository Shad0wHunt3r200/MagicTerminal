package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var dirCheckCmd = &cobra.Command{
	Use:   "dir-check",
	Short: "Check directory context",
	Run: func(cmd *cobra.Command, args []string) {
		gitCmd := exec.Command("git", "status", "--porcelain")
		_, err := gitCmd.Output()

		if err != nil {
			fmt.Println("📂 You are in a normal folder")
			return
		}

		fmt.Println("🌲 You are in a Git Repository")
	},
}

func init() {
	// This links dir-check directly to the root command automatically
	rootCmd.AddCommand(dirCheckCmd)
}