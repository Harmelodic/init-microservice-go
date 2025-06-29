package commons

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLogger_JSON(t *testing.T) {
	// Given
	var logBuffer bytes.Buffer
	logger := NewLogger(LogFormatJSON, &logBuffer)

	// When
	logger.Info("Message")

	// Then
	logMessage := logBuffer.String()
	assert.Equal(t, "{", logMessage[:1]) // Quick and dirty checking it's in JSON format
	assert.Contains(t, logMessage, "Message")
}

func TestNewLogger_Text(t *testing.T) {
	// Given
	var logBuffer bytes.Buffer
	logger := NewLogger(LogFormatTEXT, &logBuffer)

	// When
	logger.Info("Message")

	// Then
	logMessage := logBuffer.String()
	assert.Equal(t, "time", logMessage[:4]) // First key in log is time
	assert.Contains(t, logMessage, "Message")
}

func TestNewLogger_DefaultIsJSON(t *testing.T) {
	// Given
	var logBuffer bytes.Buffer
	logger := NewLogger("", &logBuffer)

	// When
	logger.Info("Message")

	// Then
	logMessage := logBuffer.String()
	assert.Equal(t, "{", logMessage[:1]) // Quick and dirty checking it's in JSON format
	assert.Contains(t, logMessage, "Message")
}

func TestNewLogger_LogIncludesSource(t *testing.T) {
	// Given
	var logBuffer bytes.Buffer
	logger := NewLogger(LogFormatTEXT, &logBuffer)

	// When
	logger.Info("Message")

	// Then
	logMessage := logBuffer.String()
	assert.Contains(t, logMessage, "commons/logging_test.go")
	assert.Contains(t, logMessage, "Message")
}
