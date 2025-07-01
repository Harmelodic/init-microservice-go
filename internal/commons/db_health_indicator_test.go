package commons

import (
	"github.com/stretchr/testify/assert"
	"log/slog"
	"testing"
)

func TestNewDbHealthIndicator_IndicateHealth(t *testing.T) {
	// Given
	database, cleanUp := NewMockDb(t)
	healthIndicator := NewDbHealthIndicator("testDb", database, slog.New(slog.DiscardHandler))
	defer cleanUp()

	// When
	name, isHealthy := healthIndicator.IndicateHealth()

	// Then
	assert.Equal(t, "testDb", name)
	assert.True(t, isHealthy)
}

func TestNewDbHealthIndicator_IndicateHealthFail(t *testing.T) {
	// Given
	database, cleanUp := NewMockDb(t)
	healthIndicator := NewDbHealthIndicator("testDb", database, slog.New(slog.DiscardHandler))
	cleanUp() // Clean up database before using it to induce connection error

	// When
	name, isHealthy := healthIndicator.IndicateHealth()

	// Then
	assert.Equal(t, "testDb", name)
	assert.False(t, isHealthy)
}
