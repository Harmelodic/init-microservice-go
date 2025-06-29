package commons

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"testing"
)

func TestNewAppDatabase_IndicateHealth(t *testing.T) {
	// Given
	appDatabase, done := NewAppDatabaseWithTestcontainers(t, "postgres", slog.New(slog.DiscardHandler))
	defer done()

	// When
	name, isHealthy := appDatabase.IndicateHealth()

	// Then
	assert.Equal(t, "TestAppDatabase", name)
	assert.True(t, isHealthy)
}

func TestNewAppDatabase_IndicateHealthFail(t *testing.T) {
	db, err := sql.Open("postgres", "postgres://postgres:password@localhost/postgres?sslmode=disable")
	if err != nil {
		t.Errorf("Failed to open database: %s", err)
	}
	defer func() {
		err = db.Close()
		if err != nil {
			t.Error("Failed to close DB")
		}
	}()

	appDatabase := AppDatabase{
		Name:   "TestAppDatabase",
		Db:     db,
		Logger: slog.New(slog.DiscardHandler),
	}

	name, isHealthy := appDatabase.IndicateHealth()
	assert.Equal(t, "TestAppDatabase", name)
	assert.False(t, isHealthy)
}
