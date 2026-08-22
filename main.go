package main

import (
	"flag"
	"fmt"
	"os/exec"
)

func main() {
	// 1. Define a "check" flag. By default, it is false.
	runCheck := flag.Bool("check", false, "Run background system state check")
	
	// 2. Parse the arguments you type after the word "magic"
	flag.Parse()

	// 3. If you DID NOT type "magic --check", do not run the git scans!
	if !*runCheck {
		fmt.Println("👋 Welcome to Magic Command Loader.")
		fmt.Println("Type 'magic --check' to scan your folder context.")
		return
	}

	// 4. This only runs if you explicitly type: magic --check
	cmd := exec.Command("git", "status", "--porcelain")
	_, err := cmd.Output()

	if err != nil {
		fmt.Println("📂 Context: Standard Windows Directory")
		return
	}

	fmt.Println("🌲 Context: Git Repository Detected")
}