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
- **Build Targets**: macOS (ARM64), Linux (x64), Windows (x64).

### 2. `wcat`

A wrapper for the PowerShell command `Get-Content` to easily display UTF-8 encoded files without character corruption issues.

- **Functionality**:
  - `wcat <file>`: Displays the content of a file using UTF-8 encoding.
    - Command: `powershell.exe -Command "Get-Content -Path '<file>' -Encoding utf8"`

- **Platform**: Windows only.

- **Source Code**: `wcat/main.go`
- **Build Target**: Windows (x64).

### 3. `gpush`

A convenience tool to stage all changes, commit them with a timestamp, and push to the remote repository in a single command.

- **Functionality**:
  - `gpush`: Executes `git add .`, then `git commit -m "YYYY-MM-DD HH:MM"`, and finally `git push`.

- **Source Code**: `gpush/main.go`
- **Build Targets**: macOS (ARM64), Linux (x64), Windows (x64).

## How to Build

You can build the tools for a specific platform by using the `go build` command and setting the `GOOS` and `GOARCH` environment variables. From within the directory of the tool you want to build (e.g., `pdoc`, `wcat`):

- **macOS (ARM64):**
  ```sh
  GOOS=darwin GOARCH=arm64 go build
  ```
- **Linux (x64):**
  ```sh
  GOOS=linux GOARCH=amd64 go build
  ```
- **Windows (x64):**
  ```sh
  GOOS=windows GOARCH=amd64 go build -o <tool_name>.exe
  ```

---

## User Preferences

- **Language**: The user prefers to communicate in Japanese (日本語).