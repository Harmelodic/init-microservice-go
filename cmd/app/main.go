package main

import (
	"log/slog"
	"os"

	"github.com/Harmelodic/init-microservice-go/internal/commons"
)

const (
	exitCodeFailedToLoadCommandLineFlags        = 1
	exitCodeFailedToCompleteDependencyInjection = 2
	exitCodeFailedToStartGinEngine              = 3
)

// main is the entrypoint to the microservice. Here we:
// 1. Trigger dependency injection (to initialise everything that needs to be initialised)
// 2. Run the resulting gin engine to start the web server.
func main() {
	logger := commons.NewLogger(commons.LogFormatJSON, os.Stdout)
	logger.Info("Starting service...")

	var appConfig *appConfig

	err := loadAppConfigFromCommandFlags(appConfig, logger)
	if err != nil {
		logger.Error("Error occurred when parsing command line flags", slog.String("error", err.Error()))
		os.Exit(exitCodeFailedToLoadCommandLineFlags)
	}

	engine, err := dependencyInjection(logger, appConfig)
	if err != nil {
		logger.Error("Error doing dependency injection to configure app", slog.String("error", err.Error()))
		os.Exit(exitCodeFailedToCompleteDependencyInjection)
	}

	logger.Info("Starting application on port 8080")

	err = engine.Run(":8080")
	if err != nil {
		logger.Error("Error occurred when starting Gin app. Exiting", slog.String("error", err.Error()))
		os.Exit(exitCodeFailedToStartGinEngine)
	}
}
