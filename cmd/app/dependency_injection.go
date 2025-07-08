package main

import (
	"github.com/Harmelodic/init-microservice-go/internal/account"
	"github.com/Harmelodic/init-microservice-go/internal/commons"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log/slog"
)

func dependencyInjection(logger *slog.Logger, appConfig *AppConfig) (*gin.Engine, error) {
	database, err := sqlx.Connect("postgres", appConfig.DbConnectionString)
	if err != nil {
		logger.Error("Failed to open postgres database",
			slog.String("datasource", appConfig.DbConnectionString),
			slog.String("error", err.Error()))
		return nil, err
	}

	// Run migrations before continuing, in case anything accesses the database as part of initialisation.
	err = commons.RunMigrations(database.DB, appConfig.MigrationsDirectory, logger)
	if err != nil {
		logger.Error("Failed to run migrations", slog.String("error", err.Error()))
		return nil, err
	}

	// TODO: Configure OpenTelemetry for tracing instrumentation

	engine := commons.NewGinEngine("init-microservice-go", logger)
	logger.Info("Gin engine configured")

	dbHealthIndicator := commons.NewDbHealthIndicator("appDb", database, logger)

	accountRepository := account.DefaultRepository{Db: database, Logger: logger}
	accountService := account.DefaultService{Repository: &accountRepository}

	account.Controller(engine, &accountService, logger)

	commons.LivenessController(engine)
	commons.ReadinessController(engine, dbHealthIndicator)

	return engine, nil
}
