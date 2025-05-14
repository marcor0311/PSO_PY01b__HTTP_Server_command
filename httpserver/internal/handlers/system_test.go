package handlers_test

import (
	"testing"
	"httpserver/internal/handlers"

)
func TestTimestamp_ReturnsValue(t *testing.T) {
	ts := handlers.Timestamp()
	if ts == "" {
		t.Error("expected non-empty timestamp, got empty string")
	}
}