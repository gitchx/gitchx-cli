# Project Context for Gemini

This project contains simple command-line interface (CLI) tools written in Go to act as wrappers for frequently used commands.

## Tools

### 1. `pdoc`

A tool for converting files between Markdown (GFM) and HTML using `pandoc`.

- **Functionality**:
  - `pdoc <file.md>`: Converts a Markdown file to an HTML file. It uses a Lua filter (`no-id.lua`) to prevent Pandoc from automatically adding IDs to headers.
    - Command: `pandoc <file.md> -f gfm -t html -o <file.html> --lua-filter=no-id.lua`
  - `pdoc <file.html>`: Converts an HTML file to a GitHub Flavored Markdown (`gfm`) file.
    - Command: `pandoc <file.html> -f html -t gfm -o <file.md>`

- **Dependencies**:
  - Requires `pandoc` to be installed and available in the system's PATH.

- **Setup**:
  - The `pdoc` executable must be in the same directory as the `no-id.lua` file.

- **Source Code**: `pdoc/main.go`
- **Lua Filter**: `pdoc/no-id.lua`
- **Build Targets**: macOS (darwin/amd64) and Windows (amd64).

### 2. `wcat`

A wrapper for the PowerShell command `Get-Content` to easily display UTF-8 encoded files without character corruption issues.

- **Functionality**:
  - `wcat <file>`: Displays the content of a file using UTF-8 encoding.
    - Command: `powershell.exe -Command "Get-Content -Path '<file>' -Encoding utf8"`

- **Platform**: Windows only.

- **Source Code**: `wcat/main.go`
- **Build Target**: Windows (amd64).

## How to Build

Use the Go compiler to build the executables. From within each tool's directory (`pdoc` or `wcat`):

- **For macOS/Linux**:
  ```sh
  go build
  ```

- **For Windows (amd64) Cross-compilation**:
  ```sh
  GOOS=windows GOARCH=amd64 go build
  ```

---

## User Preferences

- **Language**: The user prefers to communicate in Japanese (日本語).