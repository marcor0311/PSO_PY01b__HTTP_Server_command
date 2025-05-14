package handlers_test

import (
	"strings"
	"testing"
	"httpserver/internal/handlers"
)

func TestHelpText_ContainsRoutes(t *testing.T) {
	help := handlers.HelpText()
	if !strings.Contains(help, "/fibonacci") || !strings.Contains(help, "/random") {
		t.Error("HelpText() missing expected routes")
	}
}
