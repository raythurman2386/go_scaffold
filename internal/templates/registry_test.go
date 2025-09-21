package templates

import (
	"testing"
)

func TestNewTemplateRegistry(t *testing.T) {
	registry := NewTemplateRegistry()

	if registry == nil {
		t.Fatal("Expected registry to be non-nil")
	}

	if registry.templates == nil {
		t.Fatal("Expected templates map to be initialized")
	}
}

func TestGetTemplate(t *testing.T) {
	registry := NewTemplateRegistry()

	// Test existing templates
	templates := []string{"full", "cli", "api", "lib"}
	for _, name := range templates {
		template, err := registry.GetTemplate(name)
		if err != nil {
			t.Errorf("Expected no error for template '%s', got %v", name, err)
		}
		if template == nil {
			t.Errorf("Expected template '%s' to be non-nil", name)
		}
		if template.Name() != name {
			t.Errorf("Expected template name '%s', got '%s'", name, template.Name())
		}
	}

	// Test non-existing template
	_, err := registry.GetTemplate("nonexistent")
	if err == nil {
		t.Error("Expected error for non-existing template")
	}
}

func TestListTemplates(t *testing.T) {
	registry := NewTemplateRegistry()

	names := registry.ListTemplates()

	expected := []string{"full", "cli", "api", "lib"}
	if len(names) != len(expected) {
		t.Errorf("Expected %d templates, got %d", len(expected), len(names))
	}

	// Check all expected are present (order may vary)
	for _, exp := range expected {
		found := false
		for _, name := range names {
			if name == exp {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected template '%s' in list", exp)
		}
	}
}

func TestGetTemplateInfo(t *testing.T) {
	registry := NewTemplateRegistry()

	template, err := registry.GetTemplateInfo("api")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if template == nil {
		t.Error("Expected template to be non-nil")
	}
	if template.Name() != "api" {
		t.Errorf("Expected name 'api', got '%s'", template.Name())
	}
}
