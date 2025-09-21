package templates

import "fmt"

// TemplateRegistry holds all available templates
type TemplateRegistry struct {
	templates map[string]Template
}

// NewTemplateRegistry creates a new template registry
func NewTemplateRegistry() *TemplateRegistry {
	return &TemplateRegistry{
		templates: map[string]Template{
			"full": NewFullTemplate(),
			"cli":  NewCLITemplate(),
			"api":  NewAPITemplate(),
			"lib":  NewLibTemplate(),
		},
	}
}

// GetTemplate returns a template by name
func (r *TemplateRegistry) GetTemplate(name string) (Template, error) {
	template, exists := r.templates[name]
	if !exists {
		return nil, fmt.Errorf("template '%s' not found", name)
	}
	return template, nil
}

// ListTemplates returns all available template names
func (r *TemplateRegistry) ListTemplates() []string {
	names := make([]string, 0, len(r.templates))
	for name := range r.templates {
		names = append(names, name)
	}
	return names
}

// GetTemplateInfo returns information about a template
func (r *TemplateRegistry) GetTemplateInfo(name string) (Template, error) {
	return r.GetTemplate(name)
}
