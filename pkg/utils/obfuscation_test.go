package utils

import (
	"strings"
	"testing"
)

func TestExtractVariables(t *testing.T) {
	code := `
		int myVar = 10;
		String secret_key = "abc";
	`
	vars := ExtractVariables(code, "java")
	if len(vars) != 2 || vars[0] != "myVar" || vars[1] != "secret_key" {
		t.Fatalf("Failed to extract vars: %v", vars)
	}
}

func TestMangleWhitespace(t *testing.T) {
	code := `int main() { const char* s = "hello world"; return 0; }`
	mangled := MangleWhitespace(code)
	
	if !strings.Contains(mangled, "\"hello world\"") {
		t.Fatalf("Mangler corrupted the string literal: %s", mangled)
	}
	
	if mangled == code {
		t.Fatalf("Mangler did not modify whitespace")
	}
}
