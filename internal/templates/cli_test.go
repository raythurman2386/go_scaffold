package templates

import (
	"testing"
)

func TestNewCLITemplate(t *testing.T) {
	template := NewCLITemplate()

	if template.Name() != "cli" {
		t.Errorf("Expected name 'cli', got '%s'", template.Name())
	}

	if template.Description() != "Command-line interface tool" {
		t.Errorf("Expected description 'Command-line interface tool', got '%s'", template.Description())
	}

	expectedDirs := []string{
		"cmd",
		"pkg",
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
