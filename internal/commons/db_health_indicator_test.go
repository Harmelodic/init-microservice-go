package commons_test

import (
	"log/slog"
	"testing"

	"github.com/Harmelodic/init-microservice-go/internal/commons"
	"github.com/stretchr/testify/assert"
)

func TestNewDbHealthIndicator_IndicateHealth(t *testing.T) {
	t.Parallel()
	// Given
	logger := slog.New(slog.DiscardHandler)
	database, cleanUp := commons.NewMockDb(t, "../../migrations", logger)
	healthIndicator := commons.NewDbHealthIndicator("testDb", database, logger)

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
	database, cleanUp := commons.NewMockDb(t, "../../migrations", logger)
	healthIndicator := commons.NewDbHealthIndicator("testDb", database, logger)

	cleanUp() // Clean up database before using it to induce connection error

	// When
	name, isHealthy := healthIndicator.IndicateHealth()

	// Then
	assert.Equal(t, "testDb", name)
	assert.False(t, isHealthy)
}
