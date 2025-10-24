package commons_test

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Harmelodic/init-microservice-go/internal/commons"
	"github.com/stretchr/testify/assert"
)

// Mocks

type testHealthIndicator struct {
	status bool
}

func (hi testHealthIndicator) IndicateHealth() (string, bool) {
	return "indicator", hi.status
}

// Tests

func TestLivenessController_UpSolo(t *testing.T) {
	t.Parallel()

	testEngine := commons.NewGinEngine("test", slog.New(slog.DiscardHandler))
	commons.LivenessController(testEngine)

	responseRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health/liveness", http.NoBody)
	testEngine.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestLivenessController_UpWithIndicators(t *testing.T) {
	t.Parallel()

	testEngine := commons.NewGinEngine("test", slog.New(slog.DiscardHandler))
	commons.LivenessController(testEngine, testHealthIndicator{true}, testHealthIndicator{true})

	responseRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health/liveness", http.NoBody)
	testEngine.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestLivenessController_DownWhenSomeDown(t *testing.T) {
	t.Parallel()

	testEngine := commons.NewGinEngine("test", slog.New(slog.DiscardHandler))
	commons.LivenessController(testEngine, testHealthIndicator{true}, testHealthIndicator{false})

	responseRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health/liveness", http.NoBody)
	testEngine.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusServiceUnavailable, responseRecorder.Code)
}

func TestLivenessController_DownWhenAllDown(t *testing.T) {
	t.Parallel()

	testEngine := commons.NewGinEngine("test", slog.New(slog.DiscardHandler))
	commons.LivenessController(testEngine, testHealthIndicator{false}, testHealthIndicator{false})

	responseRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health/liveness", http.NoBody)
	testEngine.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusServiceUnavailable, responseRecorder.Code)
}

func TestReadinessController_UpSolo(t *testing.T) {
	t.Parallel()

	testEngine := commons.NewGinEngine("test", slog.New(slog.DiscardHandler))
	commons.ReadinessController(testEngine)

	responseRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health/readiness", http.NoBody)
	testEngine.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestReadinessController_UpWithIndicators(t *testing.T) {
	t.Parallel()

	testEngine := commons.NewGinEngine("test", slog.New(slog.DiscardHandler))
	commons.ReadinessController(testEngine, testHealthIndicator{true}, testHealthIndicator{true})

	responseRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health/readiness", http.NoBody)
	testEngine.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestReadinessController_DownWhenSomeDown(t *testing.T) {
	t.Parallel()

	testEngine := commons.NewGinEngine("test", slog.New(slog.DiscardHandler))
	commons.ReadinessController(testEngine, testHealthIndicator{true}, testHealthIndicator{false})

	responseRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health/readiness", http.NoBody)
	testEngine.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusServiceUnavailable, responseRecorder.Code)
}

func TestReadinessController_DownWhenAllDown(t *testing.T) {
	t.Parallel()

	testEngine := commons.NewGinEngine("test", slog.New(slog.DiscardHandler))
	commons.ReadinessController(testEngine, testHealthIndicator{false}, testHealthIndicator{false})

	responseRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health/readiness", http.NoBody)
	testEngine.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusServiceUnavailable, responseRecorder.Code)
}
