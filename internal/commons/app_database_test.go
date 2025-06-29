package commons

import (
	"github.com/stretchr/testify/assert"
	"log/slog"
	"testing"
)

func TestNewAppDatabase_IndicateHealth(t *testing.T) {
	// Given
	appDatabase, cleanUp := NewMockAppDatabase(t, "postgres", slog.New(slog.DiscardHandler))
	defer cleanUp()

	// When
	name, isHealthy := appDatabase.IndicateHealth()

	// Then
	assert.Equal(t, "TestAppDatabase", name)
	assert.True(t, isHealthy)
}

func TestNewAppDatabase_IndicateHealthFail(t *testing.T) {
	// Given
	appDatabase, cleanUp := NewMockAppDatabase(t, "postgres", slog.New(slog.DiscardHandler))
	cleanUp() // Clean up database before using it to induce error

	// When
	name, isHealthy := appDatabase.IndicateHealth()

	// Then
	assert.Equal(t, "TestAppDatabase", name)
	assert.False(t, isHealthy)
}
