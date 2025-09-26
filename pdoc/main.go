package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: pdoc <input_file.md | input_file.html>")
		os.Exit(1)
	}

	inputFile := os.Args[1]

	if strings.HasSuffix(inputFile, ".md") {
		// Convert Markdown to HTML
		htmlFile := strings.TrimSuffix(inputFile, ".md") + ".html"

		exePath, err := os.Executable()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting executable path: %v\n", err)
			os.Exit(1)
		}
		exeDir := filepath.Dir(exePath)
		luaFilterPath := filepath.Join(exeDir, "no-id.lua")

		if _, err := os.Stat(luaFilterPath); os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "Error: no-id.lua not found in the same directory as the pdoc executable (%s)\n", exeDir)
			os.Exit(1)
		}

		cmd := exec.Command("pandoc", inputFile, "-f", "gfm", "-t", "html", "-o", htmlFile, "--lua-filter="+luaFilterPath)

		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error converting Markdown to HTML: %v\n", err)
			fmt.Fprintf(os.Stderr, "Pandoc output: %s\n", string(output))
			os.Exit(1)
		}
		fmt.Printf("Successfully converted %s to %s\n", inputFile, htmlFile)

	} else if strings.HasSuffix(inputFile, ".html") {
		// Convert HTML to Markdown
		mdFile := strings.TrimSuffix(inputFile, ".html") + ".md"

		cmd := exec.Command("pandoc", inputFile, "-f", "html", "-t", "gfm", "-o", mdFile)

		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error converting HTML to Markdown: %v\n", err)
			fmt.Fprintf(os.Stderr, "Pandoc output: %s\n", string(output))
			os.Exit(1)
		}
		fmt.Printf("Successfully converted %s to %s\n", inputFile, mdFile)

	} else {
		fmt.Fprintln(os.Stderr, "Error: Input file must be a .md or .html file")
		os.Exit(1)
	}
}