package core

import "github.com/rs/zerolog"

type Logger struct {
}

func NewLogger() *Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	return &Logger{}
}
