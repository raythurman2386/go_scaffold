package templates

import (
	"testing"
)

func TestNewFullTemplate(t *testing.T) {
	template := NewFullTemplate()

	if template.Name() != "full" {
		t.Errorf("Expected name 'full', got '%s'", template.Name())
	}

	if template.Description() != "Full-stack web application with API, services, and templates" {
		t.Errorf("Expected description 'Full-stack web application with API, services, and templates', got '%s'", template.Description())
	}

	expectedDirs := []string{
		"cmd",
		"internal/api",
		"internal/service",
		"internal/repository",
		"internal/domain",
		"pkg",
		"configs",
		"web/templates",
		"scripts",
	}

	dirs := template.Dirs()
	if len(dirs) != len(expectedDirs) {
		t.Errorf("Expected %d directories, got %d", len(expectedDirs), len(dirs))
	}

	for i, dir := range expectedDirs {
		if i >= len(dirs) || dirs[i] != dir {
			t.Errorf("Expected dir '%s' at index %d, got '%s'", dir, i, dirs[i])
		}
	}

	files := template.Files()
	expectedFiles := []string{
		"cmd/main.go",
		".gitignore",
		"README.md",
		"configs/config.yaml",
	}

	if len(files) != len(expectedFiles) {
		t.Errorf("Expected %d files, got %d", len(expectedFiles), len(files))
	}

	for _, file := range expectedFiles {
		if _, exists := files[file]; !exists {
			t.Errorf("Expected file '%s' not found in template", file)
		}
	}
}
