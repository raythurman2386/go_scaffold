package templates

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// Template defines the interface for project templates
type Template interface {
	Name() string
	Description() string
	Dirs() []string
	Files() map[string]string
}

// BaseTemplate provides common functionality for all templates
type BaseTemplate struct {
	name        string
	description string
	dirs        []string
	files       map[string]string
}

func (t *BaseTemplate) Name() string {
	return t.name
}

func (t *BaseTemplate) Description() string {
	return t.description
}

func (t *BaseTemplate) Dirs() []string {
	return t.dirs
}

func (t *BaseTemplate) Files() map[string]string {
	return t.files
}

// CreateProject creates a new project using the template
func CreateProject(template Template, projectName string) error {
	fmt.Printf("🚀 Scaffolding new Go project: %s using template: %s\n", projectName, template.Name())

	// Create directories
	for _, dir := range template.Dirs() {
		fullPath := filepath.Join(projectName, dir)
		if err := createDir(fullPath); err != nil {
			return fmt.Errorf("failed to create directory %s: %v", fullPath, err)
		}
	}

	// Create files
	for path, content := range template.Files() {
		fullPath := filepath.Join(projectName, path)
		processedContent := processContent(content, projectName)
		if err := createFile(fullPath, processedContent); err != nil {
			return fmt.Errorf("failed to create file %s: %v", fullPath, err)
		}
	}

	// Initialize Go module
	fmt.Println("▶️  Running 'go mod init'...")
	cmd := exec.Command("go", "mod", "init", projectName)
	cmd.Dir = projectName
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to run 'go mod init': %s\n%v", string(output), err)
	}

	fmt.Println("\n✅ Project scaffolding complete!")
	fmt.Printf("📂 Your new project is ready in the '%s' directory.\n", projectName)
	fmt.Println("\nNext Steps:")
	fmt.Printf("1. Change into the new directory: cd %s\n", projectName)
	fmt.Println("2. Install dependencies: go mod tidy")

	return nil
}

// processContent replaces placeholders in template content
func processContent(content, projectName string) string {
	content = strings.ReplaceAll(content, "%s", projectName)
	content = strings.ReplaceAll(content, "%date%", time.Now().Format("2006-01-02"))
	return content
}

// createDir creates a directory if it doesn't already exist
func createDir(path string) error {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return fmt.Errorf("could not create directory %s: %w", path, err)
	}
	fmt.Printf("✔️  Created directory: %s\n", path)
	return nil
}

// createFile creates a new file with the given content
func createFile(path, content string) error {
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
