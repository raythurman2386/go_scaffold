package main

import (
	"fmt"
	"log"

	"go_scaffold/internal/templates"

	"github.com/spf13/cobra"
)

var registry = templates.NewTemplateRegistry()

func main() {
	var rootCmd = &cobra.Command{
		Use:   "go_scaffold",
		Short: "A powerful command-line tool for scaffolding new Go projects",
		Long: `Go Scaffold Tool

A powerful command-line tool for scaffolding new Go projects with multiple templates
based on popular Go project setups.

Available templates:
- full: Full-stack web application with API, services, and templates
- cli: Command-line interface tool
- api: REST API server
- lib: Go library/package`,
	}

	// Add subcommands for each template
	rootCmd.AddCommand(createTemplateCommand("full", "Create a full-stack web application"))
	rootCmd.AddCommand(createTemplateCommand("cli", "Create a command-line interface tool"))
	rootCmd.AddCommand(createTemplateCommand("api", "Create a REST API server"))
	rootCmd.AddCommand(createTemplateCommand("lib", "Create a Go library/package"))

	// Add list command
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List available templates",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Available templates:")
			for _, name := range registry.ListTemplates() {
				template, _ := registry.GetTemplateInfo(name)
				fmt.Printf("  %s: %s\n", name, template.Description())
			}
		},
	}
	rootCmd.AddCommand(listCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func createTemplateCommand(templateName, description string) *cobra.Command {
	return &cobra.Command{
		Use:   fmt.Sprintf("%s [project-name]", templateName),
		Short: description,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			projectName := args[0]

			template, err := registry.GetTemplate(templateName)
			if err != nil {
				log.Fatal(err)
			}

			if err := templates.CreateProject(template, projectName); err != nil {
				log.Fatal(err)
			}

			// Print template-specific next steps
			printNextSteps(templateName, projectName)
		},
	}
}

func printNextSteps(templateName, projectName string) {
	switch templateName {
	case "full", "api":
		fmt.Println("3. Start the server: go run cmd/main.go")
	case "cli":
		fmt.Println("3. Run the CLI: go run cmd/main.go --help")
	case "lib":
		fmt.Println("3. Use the library in your Go projects")
	}
}
