package commons_test

import (
	"bytes"
	"github.com/Harmelodic/init-microservice-go/internal/commons"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLogger_JSON(t *testing.T) {
	t.Parallel()
	// Given
	var logBuffer bytes.Buffer

	logger := commons.NewLogger(commons.LogFormatJSON, &logBuffer)

	// When
	logger.Info("Message")

	// Then
	logMessage := logBuffer.String()
	assert.Equal(t, "{", logMessage[:1]) // Quick and dirty checking it's in JSON format
	assert.Contains(t, logMessage, "Message")
}

func TestNewLogger_Text(t *testing.T) {
	t.Parallel()
	// Given
	var logBuffer bytes.Buffer

	logger := commons.NewLogger(commons.LogFormatTEXT, &logBuffer)

	// When
	logger.Info("Message")

	// Then
	logMessage := logBuffer.String()
	assert.Equal(t, "time", logMessage[:4]) // First key in log is time
	assert.Contains(t, logMessage, "Message")
}

func TestNewLogger_DefaultIsJSON(t *testing.T) {
	t.Parallel()
	// Given
	var logBuffer bytes.Buffer

	logger := commons.NewLogger("", &logBuffer)

	// When
	logger.Info("Message")

	// Then
	logMessage := logBuffer.String()
	assert.Equal(t, "{", logMessage[:1]) // Quick and dirty checking it's in JSON format
	assert.Contains(t, logMessage, "Message")
}

func TestNewLogger_LogIncludesSource(t *testing.T) {
	t.Parallel()
	// Given
	var logBuffer bytes.Buffer

	logger := commons.NewLogger(commons.LogFormatTEXT, &logBuffer)

	// When
	logger.Info("Message")

	// Then
	logMessage := logBuffer.String()
	assert.Contains(t, logMessage, "commons/logging_test.go")
	assert.Contains(t, logMessage, "Message")
}
