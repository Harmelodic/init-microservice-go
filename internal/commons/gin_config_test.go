package commons_test

import (
	"bytes"
	"github.com/Harmelodic/init-microservice-go/internal/commons"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewGinEngine_RecoversFromPanics(t *testing.T) {
	t.Parallel()
	// Given
	testEngine := commons.NewGinEngine("test", slog.New(slog.DiscardHandler))
	testEngine.GET("/endpoint", func(_ *gin.Context) {
		panic(1)
	})

	// When
	responseRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/endpoint", http.NoBody)
	testEngine.ServeHTTP(responseRecorder, req)

	// Then
	assert.Equal(t, http.StatusInternalServerError, responseRecorder.Code)
	assert.Empty(t, responseRecorder.Body.String())
}

func TestNewGinEngine_LogsConfiguredCorrectly(t *testing.T) {
	t.Parallel()
	// Given
	var logBuffer bytes.Buffer

	logger := slog.New(slog.NewTextHandler(&logBuffer, nil))
	testEngine := commons.NewGinEngine("test", logger)
	testEngine.GET("/endpoint", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{})
	})

	// When
	responseRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/endpoint", http.NoBody)
	testEngine.ServeHTTP(responseRecorder, req)

	// Then
	logOutput := logBuffer.String()
	assert.Equal(t, 1, strings.Count(logOutput, "\n")) // Only one log made for one request/response

	// Ensure important request/response info is there:
	assert.Contains(t, logOutput, "request.method=GET")
	assert.Contains(t, logOutput, "request.path=/endpoint")
	assert.Contains(t, logOutput, "response.status=200")

	// TODO when tracing instrumentation configured
	// Assert log contains trace ID for connecting logs to traces:
	// assert.Contains(t, logOutput, sloggin.TraceIDKey)
}

func TestGinReadyForProductionUse(t *testing.T) {
	t.Parallel()
	commons.NewGinEngine("test", slog.New(slog.DiscardHandler))
	assert.Equal(t, gin.ReleaseMode, gin.Mode())
}
