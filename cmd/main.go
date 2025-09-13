package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	// 1. Get the project name from the command-line arguments.
	if len(os.Args) < 2 {
		log.Fatal("Error: Please provide a project name (e.g., 'go run scaffold.go my-cloud-app')")
	}
	projectName := os.Args[1]

	fmt.Printf("🚀 Scaffolding new Go project: %s\n", projectName)

	// 2. Define the directory structure based on Go best practices.
	// This structure separates concerns and makes the project easy to navigate.
	dirs := []string{
		"cmd",
		"internal/api",        // For API handlers/controllers
		"internal/service",    // For business logic
		"internal/repository", // For database interaction
		"internal/domain",     // For your core models/structs
		"pkg",                 // For public libraries (if any)
		"configs",             // For configuration files
		"web/templates",       // For server-side HTML templates
		"scripts",             // For helper scripts (build, deploy, etc.)
	}

	for _, dir := range dirs {
		fullPath := filepath.Join(projectName, dir)
		if err := createDir(fullPath); err != nil {
			log.Fatalf("Failed to create directory %s: %v", fullPath, err)
		}
	}

	// 3. Define the initial files to create.
	files := map[string]string{
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
		"README.md": fmt.Sprintf("# %s\n\n_Project created on: %s_\n\n## Description\n\n(Your project description here)\n\n## Getting Started\n\n1.  **Install Dependencies:**\n\n    `go mod tidy`\n\n2.  **Run the application:**\n\n    `go run cmd/main.go`\n", strings.ToTitle(projectName), "2025-09-12"),
		"configs/config.yaml": `server:
  port: 8080

database:
  host: "localhost"
  port: 5432
  user: "postgres"
  password: "password"
  dbname: "yourdb"
`,
	}

	for path, content := range files {
		fullPath := filepath.Join(projectName, path)
		if err := createFile(fullPath, content); err != nil {
			log.Fatalf("Failed to create file %s: %v", fullPath, err)
		}
	}

	// 4. Initialize the Go module for the new project.
	// This is a crucial step for managing dependencies.
	fmt.Println("▶️  Running 'go mod init'...")
	cmd := exec.Command("go", "mod", "init", projectName)
	cmd.Dir = projectName // Run the command inside the new project directory
	if output, err := cmd.CombinedOutput(); err != nil {
		log.Fatalf("Failed to run 'go mod init': %s\n%v", string(output), err)
	}

	// 5. Final success message with next steps.
	fmt.Println("\n✅ Project scaffolding complete!")
	fmt.Println("📂 Your new project is ready in the '" + projectName + "' directory.")
	fmt.Println("\nNext Steps:")
	fmt.Printf("1. Change into the new directory: cd %s\n", projectName)
	fmt.Println("2. Install dependencies: go mod tidy")
	fmt.Println("3. Start the server: go run cmd/main.go")
}

// createDir creates a directory if it doesn't already exist.
func createDir(path string) error {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return fmt.Errorf("could not create directory %s: %w", path, err)
	}
	fmt.Printf("✔️  Created directory: %s\n", path)
	return nil
}

// createFile creates a new file with the given content.
func createFile(path string, content string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create file %s: %w", path, err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("could not write to file %s: %w", path, err)
	}
	fmt.Printf("✔️  Created file:      %s\n", path)
	return nil
}
