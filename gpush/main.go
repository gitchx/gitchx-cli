package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Printf("Running command: %s %v
", name, args)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to execute command '%s %v': %w", name, args, err)
	}
	return nil
}

func main() {
	// Step 1: git add .
	if err := runCommand("git", "add", "."); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Step 2: git commit
	commitMessage := time.Now().Format("2006-01-02 15:04")
	if err := runCommand("git", "commit", "-m", commitMessage); err != nil {
		fmt.Println(err)
		// It's possible there was nothing to commit. We can choose to continue to push.
		// For now, we'll exit on any error.
		os.Exit(1)
	}

	// Step 3: git push
	if err := runCommand("git", "push"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("
Successfully added, committed, and pushed changes.")
}
