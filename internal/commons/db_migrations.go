package commons

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log/slog"
)

func RunMigrations(database *sql.DB, migrationDirectory string, logger *slog.Logger) error {
	driver, err := postgres.WithInstance(database, &postgres.Config{})
	if err != nil {
		logger.Error("Failed to create driver")
		return err
	}

	migrationSource := fmt.Sprintf("file://%s", migrationDirectory)
	logger.Info("Migration source defined", slog.String("source", migrationSource))

	migrateInstance, err := migrate.NewWithDatabaseInstance(migrationSource, "postgres", driver)
	if err != nil {
		logger.Error("Failed to create migration instance")
		return err
	}

	err = migrateInstance.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			logger.Info("DB Migrations up to date")
		} else {
			logger.Error("Failed to run migrations")
			return err
		}
	}

	return nil
}
