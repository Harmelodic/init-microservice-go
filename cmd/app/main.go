package main

import (
	"github.com/Harmelodic/init-microservice-go/internal/account"
	"github.com/Harmelodic/init-microservice-go/internal/commons"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log/slog"
	"os"
)

// main is the entrypoint to the microservice. Here we:
// 1. Trigger dependency injection (to initialise everything that needs to be initialised)
// 2. Run the resulting gin engine to start the web server.
func main() {
	logger := commons.NewLogger(commons.LogFormatJSON, os.Stdout)
	logger.Info("Starting service...")

	engine := dependencyInjection(logger)

	logger.Info("Starting application on port 8080")
	err := engine.Run(":8080")
	if err != nil {
		logger.Error("Error occurred when starting Gin app. Exiting")
		os.Exit(1)
	}
}

func dependencyInjection(logger *slog.Logger) *gin.Engine {
	// TODO: Configure OpenTelemetry for tracing instrumentation

	engine := commons.NewGinEngine("init-microservice-go", logger)
	logger.Info("Gin engine configured")

	// TODO: Replace with call to service app_database, when Flyway (or alt) configured
	driver, dataSource := "postgres", "postgres://postgres:password@localhost/postgres?sslmode=disable"
	database, err := commons.NewAppDatabase("App Database", driver, dataSource, logger)
	if err != nil {
		os.Exit(1)
	}

	accountRepository := account.DefaultRepository{Db: database.Db, Logger: logger}
	accountService := account.DefaultService{Repository: &accountRepository}

	account.Controller(engine, &accountService)

	commons.LivenessController(engine)
	commons.ReadinessController(engine, database)

	return engine
}
