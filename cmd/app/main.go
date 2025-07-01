package main

import (
	"github.com/Harmelodic/init-microservice-go/internal/account"
	"github.com/Harmelodic/init-microservice-go/internal/commons"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log/slog"
	"os"
	"strings"
)

// main is the entrypoint to the microservice. Here we:
// 1. Trigger dependency injection (to initialise everything that needs to be initialised)
// 2. Run the resulting gin engine to start the web server.
func main() {
	logger := commons.NewLogger(commons.LogFormatJSON, os.Stdout)
	logger.Info("Starting service...")

	// TODO: Replace with CLI flag, injected into AppConfig
	logger.Info("Arguments given", slog.String("args", strings.Join(os.Args, " ")))
	dbMigrationsDirectory := "/home/matt/harmelodic/init-microservice-go/migrations"

	engine := dependencyInjection(logger, dbMigrationsDirectory)

	logger.Info("Starting application on port 8080")
	err := engine.Run(":8080")
	if err != nil {
		logger.Error("Error occurred when starting Gin app. Exiting")
		os.Exit(1)
	}
}

func dependencyInjection(logger *slog.Logger, dbMigrationsDirectory string) *gin.Engine {
	// TODO: Replace with datasource connection string CLI flag, injected into AppConfig
	driver, dataSource := "postgres", "postgres://init-microservice-go:password@localhost:5432/service_db?sslmode=disable"
	database, err := sqlx.Connect(driver, dataSource)
	if err != nil {
		logger.Error("Failed to open database",
			slog.String("driver", driver),
			slog.String("datasource", dataSource),
			slog.String("error", err.Error()))
		os.Exit(1)
	}

	// Run migrations before continuing, in case anything accesses the database as part of initialisation.
	err = commons.RunMigrations(database.DB, dbMigrationsDirectory, logger)
	if err != nil {
		logger.Error("Failed to run migrations", slog.String("error", err.Error()))
		os.Exit(1)
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

	return engine
}
