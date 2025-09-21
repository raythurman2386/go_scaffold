package templates

// CLITemplate represents the command-line interface tool template
type CLITemplate struct {
	BaseTemplate
}

func NewCLITemplate() *CLITemplate {
	return &CLITemplate{
		BaseTemplate: BaseTemplate{
			name:        "cli",
			description: "Command-line interface tool",
			dirs: []string{
				"cmd",
				"pkg",
			},
			files: map[string]string{
				"cmd/main.go": `package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: mycli <command>")
		os.Exit(1)
	}

	command := os.Args[1]
	switch command {
	case "hello":
		fmt.Println("Hello, World!")
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}
`,
				".gitignore": `# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with 'go test -c'
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Go workspace file
go.work
`,
				"README.md": "# %s\n\n_Project created on: %date%_\n\n## Description\n\nA command-line tool.\n\n## Usage\n\n`go run cmd/main.go hello`\n",
			},
		},
	}
}
