package commons

import (
	"github.com/stretchr/testify/assert"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testHealthIndicator struct {
	status bool
}

func (hi testHealthIndicator) IndicateHealth() (string, bool) {
	return "indicator", hi.status
}

func TestLivenessController_UpSolo(t *testing.T) {
	testEngine := NewGinEngine("test", slog.New(slog.DiscardHandler))
	LivenessController(testEngine)

	responseRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health/liveness", http.NoBody)
	testEngine.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestLivenessController_UpWithIndicators(t *testing.T) {
	testEngine := NewGinEngine("test", slog.New(slog.DiscardHandler))
	LivenessController(testEngine, testHealthIndicator{true}, testHealthIndicator{true})

	responseRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health/liveness", http.NoBody)
	testEngine.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestLivenessController_DownWhenSomeDown(t *testing.T) {
	testEngine := NewGinEngine("test", slog.New(slog.DiscardHandler))
	LivenessController(testEngine, testHealthIndicator{true}, testHealthIndicator{false})

	responseRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health/liveness", http.NoBody)
	testEngine.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusServiceUnavailable, responseRecorder.Code)
}

func TestLivenessController_DownWhenAllDown(t *testing.T) {
	testEngine := NewGinEngine("test", slog.New(slog.DiscardHandler))
	LivenessController(testEngine, testHealthIndicator{false}, testHealthIndicator{false})

	responseRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health/liveness", http.NoBody)
	testEngine.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusServiceUnavailable, responseRecorder.Code)

}

func TestReadinessController_UpSolo(t *testing.T) {
	testEngine := NewGinEngine("test", slog.New(slog.DiscardHandler))
	ReadinessController(testEngine)

	responseRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health/readiness", http.NoBody)
	testEngine.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)

}

func TestReadinessController_UpWithIndicators(t *testing.T) {
	testEngine := NewGinEngine("test", slog.New(slog.DiscardHandler))
	ReadinessController(testEngine, testHealthIndicator{true}, testHealthIndicator{true})

	responseRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health/readiness", http.NoBody)
	testEngine.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)

}

func TestReadinessController_DownWhenSomeDown(t *testing.T) {
	testEngine := NewGinEngine("test", slog.New(slog.DiscardHandler))
	ReadinessController(testEngine, testHealthIndicator{true}, testHealthIndicator{false})

	responseRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health/readiness", http.NoBody)
	testEngine.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusServiceUnavailable, responseRecorder.Code)

}

func TestReadinessController_DownWhenAllDown(t *testing.T) {
	testEngine := NewGinEngine("test", slog.New(slog.DiscardHandler))
	ReadinessController(testEngine, testHealthIndicator{false}, testHealthIndicator{false})

	responseRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health/readiness", http.NoBody)
	testEngine.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusServiceUnavailable, responseRecorder.Code)
}
