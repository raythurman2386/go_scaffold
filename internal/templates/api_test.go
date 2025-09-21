package templates

import (
	"testing"
)

func TestNewAPITemplate(t *testing.T) {
	template := NewAPITemplate()

	if template.Name() != "api" {
		t.Errorf("Expected name 'api', got '%s'", template.Name())
	}

	if template.Description() != "REST API server" {
		t.Errorf("Expected description 'REST API server', got '%s'", template.Description())
	}

	expectedDirs := []string{
		"cmd",
		"internal/api",
		"internal/service",
		"internal/repository",
		"internal/domain",
		"pkg",
		"configs",
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
		"internal/api/handlers.go",
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
