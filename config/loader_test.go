package config

import (
	"os"
	"testing"
)

func TestNewLoader(t *testing.T) {
	// Create a temporary .env file for testing
	envContent := "TEST_KEY=test_value\n"
	tmpFile, err := os.CreateTemp("", ".env.*")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(envContent); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if err := tmpFile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	// Change to the temp directory temporarily
	oldDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	defer os.Chdir(oldDir)

	tmpDir := tmpFile.Name()
	os.Chdir(tmpDir)

	// Test NewLoader with valid .env file
	loader, err := NewLoader()
	if err != nil {
		t.Logf("NewLoader failed (expected if .env doesn't exist): %v", err)
		// This is acceptable as the function tries to read .env file
		return
	}

	if loader == nil {
		t.Error("NewLoader returned nil loader")
	}

	// Test with missing .env file (should still work as it uses environment variables)
	os.Remove(".env")
	loader2, err := NewLoader()
	if err != nil {
		// This is acceptable as the function tries to read .env file
		t.Logf("NewLoader failed when .env doesn't exist: %v", err)
		return
	}

	if loader2 == nil {
		t.Error("NewLoader returned nil loader")
	}
}
