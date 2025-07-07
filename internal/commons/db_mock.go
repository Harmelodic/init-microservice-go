package commons

import (
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"log/slog"
	"testing"
)

func NewMockDb(t *testing.T, migrationsDirectory string, logger *slog.Logger) (db *sqlx.DB, cleanUp func()) {
	ctx := context.Background()

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
	cleanUp = func() {
		if err := testcontainers.TerminateContainer(postgresContainer); err != nil {
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
