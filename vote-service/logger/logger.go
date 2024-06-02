package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

// NewConsoleLogger creates a new zerolog console logger
// This function sets up a zerolog logger that writes to the console (stdout)
// with timestamps formatted according to RFC3339.
func NewConsoleLogger() zerolog.Logger {
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,         // Output destination is set to standard output (console)
		TimeFormat: time.RFC3339,      // Time format is set to RFC3339
	}
	return zerolog.New(output).With().Timestamp().Logger() // Create and return a new logger with timestamps
}
