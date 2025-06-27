package commons

import (
	"github.com/stretchr/testify/assert"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

type healthIndicator struct {
	status bool
}

func (hi healthIndicator) IndicateHealth() (string, bool) {
	return "indicator", hi.status
}

func TestLivenessController_UpSolo(t *testing.T) {
	testEngine := NewGinEngine(slog.New(slog.DiscardHandler))
	LivenessController(testEngine)

	responseRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health/liveness", http.NoBody)
	testEngine.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestLivenessController_UpWithIndicators(t *testing.T) {
	testEngine := NewGinEngine(slog.New(slog.DiscardHandler))
	LivenessController(testEngine, healthIndicator{true}, healthIndicator{true})

	responseRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health/liveness", http.NoBody)
	testEngine.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestLivenessController_DownWhenSomeDown(t *testing.T) {
	testEngine := NewGinEngine(slog.New(slog.DiscardHandler))
	LivenessController(testEngine, healthIndicator{true}, healthIndicator{false})

	responseRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health/liveness", http.NoBody)
	testEngine.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusServiceUnavailable, responseRecorder.Code)
}

func TestLivenessController_DownWhenAllDown(t *testing.T) {
	testEngine := NewGinEngine(slog.New(slog.DiscardHandler))
	LivenessController(testEngine, healthIndicator{false}, healthIndicator{false})

	responseRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health/liveness", http.NoBody)
	testEngine.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusServiceUnavailable, responseRecorder.Code)

}

func TestReadinessController_UpSolo(t *testing.T) {
	testEngine := NewGinEngine(slog.New(slog.DiscardHandler))
	ReadinessController(testEngine)

	responseRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health/readiness", http.NoBody)
	testEngine.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)

}

func TestReadinessController_UpWithIndicators(t *testing.T) {
	testEngine := NewGinEngine(slog.New(slog.DiscardHandler))
	ReadinessController(testEngine, healthIndicator{true}, healthIndicator{true})

	responseRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health/readiness", http.NoBody)
	testEngine.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)

}

func TestReadinessController_DownWhenSomeDown(t *testing.T) {
	testEngine := NewGinEngine(slog.New(slog.DiscardHandler))
	ReadinessController(testEngine, healthIndicator{true}, healthIndicator{false})

	responseRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health/readiness", http.NoBody)
	testEngine.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusServiceUnavailable, responseRecorder.Code)

}

func TestReadinessController_DownWhenAllDown(t *testing.T) {
	testEngine := NewGinEngine(slog.New(slog.DiscardHandler))
	ReadinessController(testEngine, healthIndicator{false}, healthIndicator{false})

	responseRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health/readiness", http.NoBody)
	testEngine.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusServiceUnavailable, responseRecorder.Code)
}
