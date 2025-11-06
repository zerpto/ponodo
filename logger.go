package ponodo

import "github.com/rs/zerolog"

// Logger represents the application logger instance.
// It wraps the zerolog logger and provides structured logging capabilities
// for the application.
type Logger struct {
}

// NewLogger creates and initializes a new logger instance.
// It configures the zerolog time format to use Unix timestamps and
// returns a ready-to-use logger instance.
func NewLogger() *Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	return &Logger{}
}
