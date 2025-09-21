package templates

// FullTemplate represents the full-stack web application template
type FullTemplate struct {
	BaseTemplate
}

func NewFullTemplate() *FullTemplate {
	return &FullTemplate{
		BaseTemplate: BaseTemplate{
			name:        "full",
			description: "Full-stack web application with API, services, and templates",
			dirs: []string{
				"cmd",
				"internal/api",
				"internal/service",
				"internal/repository",
				"internal/domain",
				"pkg",
				"configs",
				"web/templates",
				"scripts",
			},
			files: map[string]string{
				"cmd/main.go": `package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server is starting on :8080...")
	// Example of a simple handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
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

# Dependency directories (remove the comment below to include it)
# vendor/

# Go workspace file
go.work
`,
				"README.md": "# %s\n\n_Project created on: %date%_\n\n## Description\n\n(Your project description here)\n\n## Getting Started\n\n1.  **Install Dependencies:**\n\n    `go mod tidy`\n\n2.  **Run the application:**\n\n    `go run cmd/main.go`\n",
				"configs/config.yaml": `server:
  port: 8080

database:
  host: "localhost"
  port: 5432
  user: "postgres"
  password: "password"
  dbname: "yourdb"
`,
			},
		},
	}
}
