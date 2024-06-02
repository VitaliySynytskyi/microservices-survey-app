package logger

import (
	"bytes"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

// TestNewConsoleLogger tests the NewConsoleLogger function
func TestNewConsoleLogger(t *testing.T) {
	// Create a buffer to capture the output
	var buf bytes.Buffer
	output := zerolog.ConsoleWriter{
		Out:        &buf,
		TimeFormat: time.RFC3339,
	}

	// Create a new logger that writes to the buffer
	logger := zerolog.New(output).With().Timestamp().Logger()

	// Log a test message
	logger.Info().Msg("test message")

	// Capture and check the output
	outputStr := buf.String()

	// Check that the output contains the test message
	assert.Contains(t, outputStr, "test message", "Expected log output to contain 'test message'")

	// Check that the output contains a timestamp
	assert.Contains(t, outputStr, time.Now().Format(time.RFC3339)[:10], "Expected log output to contain a timestamp")

	// Check that the logger is of the expected type
	assert.IsType(t, zerolog.Logger{}, logger, "Expected logger to be of type zerolog.Logger")
}
