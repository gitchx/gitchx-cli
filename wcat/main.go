package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: wcat <file_path>")
		os.Exit(1)
	}

	filePath := os.Args[1]

	// Execute PowerShell command to get content with UTF-8 encoding
	cmd := exec.Command("powershell.exe", "-Command", fmt.Sprintf("Get-Content -Path '%s' -Encoding utf8", filePath))

	// Pipe the output of the command to the standard output and standard error of this program
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		os.Exit(1)
	}
}
