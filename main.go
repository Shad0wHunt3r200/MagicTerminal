package main

import (
	"flag"
	"fmt"
	"os/exec"
)

func main() {

	runHelp := flag.Bool("help", false, "Display help information")
	runDirCheck := flag.Bool("dir-check", false, "Check directory context")


	flag.Parse()

	if *runHelp {
		if *runHelp {
			fmt.Println(`
Help Information:
	
Usage: magic [options]
	
Options:
  --help         Display help information
  --dir-check    Check directory context`)
			return
		}	
		return
	}

	if *runDirCheck {
		cmd := exec.Command("git", "status", "--porcelain")
		_, err := cmd.Output()

		if err != nil {
			fmt.Println("📂 Context: Standard Directory")
			return
		}

		fmt.Println("🌲 Context: Git Repository")
		return // Stop the application here
	}

	fmt.Println("👋 Welcome to MagicTerminal.")
	fmt.Println("Type 'magic --help' to see available options.")
}