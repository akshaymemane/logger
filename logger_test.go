package logger

import (
	"bytes"
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	// Set up environment variables for testing
	os.Setenv("LOG_FILE", "test.log")
	os.Setenv("LOG_LEVEL", "debug")
	os.Setenv("LOG_FORMAT", "date|time|shortfile")

	// Create a new logger instance
	logger := New()

	// Redirect logger output for testing
	var buf bytes.Buffer
	logger.SetOutput(&buf)

	// Log messages at different log levels
	logger.Debug("This is a debug message")
	logger.Info("This is an info message")
	logger.Warning("This is a warning message")
	logger.Error("This is an error message")

	// Check if the logged messages contain the expected prefixes
	expected := []string{
		"[DEBUG]",
		"[INFO]",
		"[WARNING]",
		"[ERROR]",
	}

	loggedMessages := buf.String()
	for _, prefix := range expected {
		if !bytes.Contains([]byte(loggedMessages), []byte(prefix)) {
			t.Errorf("Expected prefix %s not found in logged message: %s", prefix, loggedMessages)
		}
	}

	// Clean up: reset environment variables
	os.Unsetenv("LOG_FILE")
	os.Unsetenv("LOG_LEVEL")
	os.Unsetenv("LOG_FORMAT")
}
