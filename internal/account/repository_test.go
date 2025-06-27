package account

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"testing"
)

// Mocks
func usePostgresContainerDb(t *testing.T) (*sql.DB, func()) {
	ctx := context.Background()

	dbName := "postgres"
	dbUser := "postgres"
	dbPassword := "password"

	t.Log("Starting container...")
	postgresContainer, err := postgres.Run(ctx,
		"postgres:latest",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		postgres.BasicWaitStrategies(),
	)
	if err != nil {
		t.Errorf("Failed to start container: %s", err)
	}
	done := func() {
		if err := testcontainers.TerminateContainer(postgresContainer); err != nil {
			t.Logf("Failed to terminate container: %s", err.Error())
		}
	}

	connectionString, err := postgresContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Errorf("Failed to build connection string: %s", err.Error())
	}
	t.Logf("Connection string established: %s", connectionString)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		t.Errorf("Failed to open database: %s", err)
	}

	return db, done
}

// Tests

func TestDefaultRepository_GetAllAccounts(t *testing.T) {
	// TODO
}

func TestDefaultRepository_GetAllAccountsError(t *testing.T) {
	// TODO
}

func TestDefaultRepository_IsHealthy(t *testing.T) {
	db, done := usePostgresContainerDb(t)
	defer done()

	repo := DefaultRepository{Db: db}

	name, isHealthy := repo.IndicateHealth()
	assert.Equal(t, "AccountRepository", name)
	assert.True(t, isHealthy)
}

func TestDefaultRepository_IsHealthyFail(t *testing.T) {
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

	repo := DefaultRepository{Db: db}

	_, isHealthy := repo.IndicateHealth()
	assert.False(t, isHealthy)
}
