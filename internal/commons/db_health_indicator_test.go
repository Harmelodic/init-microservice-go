package commons

import (
	"github.com/stretchr/testify/assert"
	"log/slog"
	"testing"
)

func TestNewDbHealthIndicator_IndicateHealth(t *testing.T) {
	t.Parallel()
	// Given
	logger := slog.New(slog.DiscardHandler)
	database, cleanUp := NewMockDb(t, "../../migrations", logger)
	healthIndicator := NewDbHealthIndicator("testDb", database, logger)
	defer cleanUp()

	// When
	name, isHealthy := healthIndicator.IndicateHealth()

	// Then
	assert.Equal(t, "testDb", name)
	assert.True(t, isHealthy)
}

func TestNewDbHealthIndicator_IndicateHealthFail(t *testing.T) {
	t.Parallel()
	// Given
	logger := slog.New(slog.DiscardHandler)
	database, cleanUp := NewMockDb(t, "../../migrations", logger)
	healthIndicator := NewDbHealthIndicator("testDb", database, logger)
	cleanUp() // Clean up database before using it to induce connection error

	// When
	name, isHealthy := healthIndicator.IndicateHealth()

	// Then
	assert.Equal(t, "testDb", name)
	assert.False(t, isHealthy)
}
