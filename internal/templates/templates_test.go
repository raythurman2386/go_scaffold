package templates

import (
	"strings"
	"testing"
	"time"
)

func TestProcessContent(t *testing.T) {
	content := "Project: %s\nDate: %date%\nModule: %s"
	projectName := "example.com/myproject"

	processed := processContent(content, projectName)

	if !strings.Contains(processed, "Project: example.com/myproject") {
		t.Errorf("Expected project name replacement, got: %s", processed)
	}

	if !strings.Contains(processed, "Module: example.com/myproject") {
		t.Errorf("Expected module name replacement, got: %s", processed)
	}

	// Check date is replaced with today's date format
	today := time.Now().Format("2006-01-02")
	if !strings.Contains(processed, "Date: "+today) {
		t.Errorf("Expected date replacement with %s, got: %s", today, processed)
	}
}
