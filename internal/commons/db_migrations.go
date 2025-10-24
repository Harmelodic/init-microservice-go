package commons

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // Pull in `file://` driver for migrations
	_ "github.com/lib/pq"                                // Pull in Postgres driver for access a Postgres DB.
)

// RunMigrations runs the DB migrations found in the given migration directory on the given sql.DB.
func RunMigrations(database *sql.DB, migrationDirectory string, logger *slog.Logger) error {
	//exhaustruct:ignore - Safe to define postgres.Config non-exhaustively.
	driver, err := postgres.WithInstance(database, &postgres.Config{})
	if err != nil {
		logger.Error("Failed to create driver")

		return fmt.Errorf("failed to create postgres driver to run migrations: %w", err)
	}

	migrationSource := "file://" + migrationDirectory
	logger.Info("Migration source defined", slog.String("source", migrationSource))

	migrateInstance, err := migrate.NewWithDatabaseInstance(migrationSource, "postgres", driver)
	if err != nil {
		logger.Error("Failed to create migration instance")

		return fmt.Errorf("failed to create migration instance to run migrations: %w", err)
	}

	err = migrateInstance.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			logger.Info("DB Migrations up to date")
		} else {
			logger.Error("Failed to run migrations")

			return fmt.Errorf("failed to run migrations: %w", err)
		}
	}

	return nil
}
