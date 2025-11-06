package ponodo

import (
	"testing"
)

func TestNewLogger(t *testing.T) {
	logger := NewLogger()

	if logger == nil {
		t.Error("NewLogger returned nil")
	}

	// Verify it returns a Logger instance
	if _, ok := interface{}(logger).(*Logger); !ok {
		t.Error("NewLogger did not return a *Logger instance")
	}
}
