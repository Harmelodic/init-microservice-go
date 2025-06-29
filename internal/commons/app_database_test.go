package commons

import (
	"github.com/stretchr/testify/assert"
	"log/slog"
	"testing"
)

func TestNewAppDatabase_IndicateHealth(t *testing.T) {
	// Given
	appDatabase, done := NewMockAppDatabase(t, "postgres", slog.New(slog.DiscardHandler))
	defer done()

	// When
	name, isHealthy := appDatabase.IndicateHealth()

	// Then
	assert.Equal(t, "TestAppDatabase", name)
	assert.True(t, isHealthy)
}

func TestNewAppDatabase_IndicateHealthFail(t *testing.T) {
	// Given
	appDatabase, done := NewMockAppDatabase(t, "postgres", slog.New(slog.DiscardHandler))
	done() // Clean up database to induce error

	// When
	name, isHealthy := appDatabase.IndicateHealth()

	// Then
	assert.Equal(t, "TestAppDatabase", name)
	assert.False(t, isHealthy)
}
