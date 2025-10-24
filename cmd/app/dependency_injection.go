package main

import (
	"fmt"
	"log/slog"

	"github.com/Harmelodic/init-microservice-go/internal/account"
	"github.com/Harmelodic/init-microservice-go/internal/commons"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func dependencyInjection(logger *slog.Logger, appConfig *appConfig) (*gin.Engine, error) {
	database, err := sqlx.Connect("postgres", appConfig.DbConnectionString)
	if err != nil {
		logger.Error("Failed to open postgres database",
			slog.String("datasource", appConfig.DbConnectionString),
			slog.String("error", err.Error()))

		return nil, fmt.Errorf("failed to open postgres database: %w", err)
	}

	// Run migrations before continuing, in case anything accesses the database as part of initialisation.
	err = commons.RunMigrations(database.DB, appConfig.MigrationsDirectory, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	engine := commons.NewGinEngine("init-microservice-go", logger)
	logger.Info("Gin engine configured")

	dbHealthIndicator := commons.NewDbHealthIndicator("appDb", database, logger)

	accountRepository := account.DefaultRepository{Db: database}
	accountService := account.DefaultService{Repository: &accountRepository}

	account.Controller(engine, &accountService, logger)

	commons.LivenessController(engine)
	commons.ReadinessController(engine, dbHealthIndicator)

	return engine, nil
}
