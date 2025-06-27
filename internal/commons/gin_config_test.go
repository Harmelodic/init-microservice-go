package commons

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Testing

func TestNewGinEngine_RecoversFromPanics(t *testing.T) {
	// Given
	testEngine := NewGinEngine(slog.New(slog.DiscardHandler))
	testEngine.GET("/endpoint", func(context *gin.Context) {
		panic(1)
	})

	// When
	responseRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/endpoint", http.NoBody)
	testEngine.ServeHTTP(responseRecorder, req)

	// Then
	assert.Equal(t, http.StatusInternalServerError, responseRecorder.Code)
	assert.Equal(t, "", responseRecorder.Body.String())
}

// TODO: Test logs contain Trace IDs
// TODO: Test Gin is prepped for production use (Release mode)
