package utils

import (
	"testing"
)

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "camelCase",
			input:    "camelCase",
			expected: "camel_case",
		},
		{
			name:     "PascalCase",
			input:    "PascalCase",
			expected: "pascal_case",
		},
		{
			name:     "already snake_case",
			input:    "snake_case",
			expected: "snake_case",
		},
		{
			name:     "multiple words",
			input:    "MyVariableName",
			expected: "my_variable_name",
		},
		{
			name:     "with numbers",
			input:    "userID123",
			expected: "user_id123",
		},
		{
			name:     "all uppercase",
			input:    "UPPERCASE",
			expected: "uppercase",
		},
		{
			name:     "all lowercase",
			input:    "lowercase",
			expected: "lowercase",
		},
		{
			name:     "single character",
			input:    "A",
			expected: "a",
		},
		{
			name:     "complex case",
			input:    "HTTPRequestHandler",
			expected: "httprequest_handler",
		},
		{
			name:     "with acronym in middle",
			input:    "XMLParser",
			expected: "xmlparser",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToSnakeCase(tt.input)
			if result != tt.expected {
				t.Errorf("ToSnakeCase(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}
