package utils

import (
	"regexp"
	"strings"
)

// ToSnakeCase converts a camelCase or PascalCase string to snake_case.
// It inserts underscores before uppercase letters that follow lowercase
// letters or numbers, then converts the entire string to lowercase.
func ToSnakeCase(str string) string {
	// Handle empty string
	if str == "" {
		return str
	}

	// Insert underscore before uppercase letters that follow lowercase letters or numbers
	re := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	snake := re.ReplaceAllString(str, `${1}_${2}`)

	// Convert to lowercase
	return strings.ToLower(snake)
}
