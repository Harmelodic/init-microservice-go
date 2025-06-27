package main

import (
	"database/sql"
	"github.com/Harmelodic/init-microservice-go/internal/account"
	"github.com/Harmelodic/init-microservice-go/internal/commons"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log/slog"
	"os"
)

// main is the entrypoint to the microservice. Here we:
// 1. Configure the gin engine
// 2. Trigger dependency injection
// 3. Run the gin engine to start the web server.
func main() {
	logger := commons.NewLogger()
	logger.Info("Starting service...")

	engine := commons.NewGinEngine(logger)

	dependencyInjection(engine, logger)

	logger.Info("Starting application on port 8080")
	err := engine.Run(":8080")
	if err != nil {
		logger.Error("Error occurred when starting Gin app. Exiting")
		os.Exit(1)
	}
}

func dependencyInjection(engine *gin.Engine, logger *slog.Logger) {
	// TODO: Replace with call to service database
	driver, dataSource := "postgres", "postgres://postgres:password@localhost/postgres?sslmode=disable"
	database, err := sql.Open(driver, dataSource)
	if err != nil {
		logger.Error(
			"Failed to open database",
			slog.String("driver", driver),
			slog.String("datasource", dataSource),
			slog.String("error", err.Error()))
		os.Exit(1)
	}

	accountRepository := account.DefaultRepository{Db: database, Logger: logger}
	accountService := account.DefaultService{Repository: &accountRepository}

	account.Controller(engine, &accountService)

	commons.LivenessController(engine)
	commons.ReadinessController(engine, &accountRepository)
}
