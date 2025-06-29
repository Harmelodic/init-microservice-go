package commons

import (
	"context"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"log/slog"
	"testing"
)

func NewAppDatabaseWithTestcontainers(t *testing.T, dbName string, logger *slog.Logger) (db *AppDatabase, done func()) {
	ctx := context.Background()

	t.Log("Starting container...")
	postgresContainer, err := postgres.Run(ctx,
		"postgres:latest",
		postgres.WithDatabase(dbName),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("password"),
		postgres.BasicWaitStrategies(),
	)
	if err != nil {
		t.Errorf("Failed to start container: %s", err)
	}
	done = func() {
		if err := testcontainers.TerminateContainer(postgresContainer); err != nil {
			t.Logf("Failed to terminate container: %s", err.Error())
		}
	}

	connectionString, err := postgresContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		done()
		t.Errorf("Failed to build connection string: %s", err.Error())
	}
	t.Logf("Connection string established: %s", connectionString)

	appDatabase, err := NewAppDatabase("TestAppDatabase", "postgres", connectionString, logger)
	if err != nil {
		done()
		t.Errorf("Failed to open database: %s", err)
	}

	return appDatabase, done
}
