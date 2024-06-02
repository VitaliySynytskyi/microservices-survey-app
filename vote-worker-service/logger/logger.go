package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

// NewConsoleLogger creates a new zerolog console logger
// This function sets up a zerolog logger that outputs to the console
func NewConsoleLogger() zerolog.Logger {
	// Configure zerolog to write to the console with a specific time format
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

	// Create a new zerolog logger with timestamp and return it
	return zerolog.New(output).With().Timestamp().Logger()
}
