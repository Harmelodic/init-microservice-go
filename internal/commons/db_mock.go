package commons

import (
	"log/slog"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Pull in postgres driver for connecting to mock DB
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

// NewMockDb is a factory method for creating a mock database instance, using Testcontainers.
//
// NOTE: Only to be used for testing (or local development).
func NewMockDb(t *testing.T, migrationsDirectory string, logger *slog.Logger) (*sqlx.DB, func()) {
	t.Helper()

	ctx := t.Context()

	t.Log("Starting container...")

	postgresContainer, err := postgres.Run(ctx,
		"postgres:latest",
		postgres.WithDatabase("mock_db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("password"),
		postgres.BasicWaitStrategies(),
	)
	if err != nil {
		t.Fatalf("Failed to start container: %s", err)
	}

	cleanUp := func() {
		err := testcontainers.TerminateContainer(postgresContainer)
		if err != nil {
			t.Fatalf("Failed to terminate container: %s", err.Error())
		}
	}

	connectionString, err := postgresContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		cleanUp()
		t.Fatalf("Failed to build connection string: %s", err.Error())
	}

	t.Logf("Connection string established: %s", connectionString)

	database, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		cleanUp()
		t.Fatalf("Failed to open database: %s", err.Error())
	}

	err = RunMigrations(database.DB, migrationsDirectory, logger)
	if err != nil {
		cleanUp()
		t.Fatalf("Failed to run migrations: %s", err.Error())
	}

	return database, cleanUp
}
