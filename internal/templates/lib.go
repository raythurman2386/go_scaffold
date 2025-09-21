package templates

// LibTemplate represents the Go library/package template
type LibTemplate struct {
	BaseTemplate
}

func NewLibTemplate() *LibTemplate {
	return &LibTemplate{
		BaseTemplate: BaseTemplate{
			name:        "lib",
			description: "Go library/package",
			dirs: []string{
				"pkg",
			},
			files: map[string]string{
				"pkg/example.go": `package pkg

// ExampleFunction is an example function
func ExampleFunction() string {
	return "Hello from the library!"
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
				"README.md": "# %s\n\n_Project created on: %date%_\n\n## Description\n\nA Go library.\n\n## Usage\n\nImport this package in your Go project.\n",
			},
		},
	}
}
