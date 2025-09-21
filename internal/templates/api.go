package templates

// APITemplate represents the REST API server template
type APITemplate struct {
	BaseTemplate
}

func NewAPITemplate() *APITemplate {
	return &APITemplate{
		BaseTemplate: BaseTemplate{
			name:        "api",
			description: "REST API server",
			dirs: []string{
				"cmd",
				"internal/api",
				"internal/service",
				"internal/repository",
				"internal/domain",
				"pkg",
				"configs",
			},
			files: map[string]string{
				"cmd/main.go": `package main

import (
	"fmt"
	"log"
	"net/http"

	"%s/internal/api"
)

func main() {
	fmt.Println("API Server is starting on :8080...")

	// Setup routes
	http.HandleFunc("/api/health", api.HealthCheck)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %%s\n", err)
	}
}
`,
				"internal/api/handlers.go": `package api

import (
	"encoding/json"
	"net/http"
)

// HealthCheck handles health check requests
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "ok"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
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
				"README.md": "# %s\n\n_Project created on: %date%_\n\n## Description\n\nA REST API server.\n\n## Getting Started\n\n1. Install dependencies: `go mod tidy`\n2. Run the server: `go run cmd/main.go`\n",
				"configs/config.yaml": `server:
  port: 8080
`,
			},
		},
	}
}
