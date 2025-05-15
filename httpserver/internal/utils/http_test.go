package utils

import (
	"testing"
)

// ParseRequestLine 

func TestParseRequestLine_Valid(t *testing.T) {
	line := "GET /hello HTTP/1.0"
	method, path, version, ok := ParseRequestLine(line)

	if !ok || method != "GET" || path != "/hello" || version != "HTTP/1.0" {
		t.Errorf("Expected valid parse, got: %v %v %v %v", method, path, version, ok)
	}
}

func TestParseRequestLine_Invalid(t *testing.T) {
	line := "INVALIDREQUEST"
	_, _, _, ok := ParseRequestLine(line)
	if ok {
		t.Error("Expected parse fail on invalid request line")
	}
}

// ExtractQuery 

func TestExtractQuery_Valid(t *testing.T) {
	queryStr := "/test?name=foo&age=42"
	values, err := ExtractQuery(queryStr)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if values.Get("name") != "foo" || values.Get("age") != "42" {
		t.Errorf("Unexpected query values: %v", values)
	}
}

func TestExtractQuery_Invalid(t *testing.T) {
	_, err := ExtractQuery("/test-without-query")
	if err == nil {
		t.Error("Expected error for missing query")
	}
}

