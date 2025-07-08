package main

import (
	"github.com/Harmelodic/init-microservice-go/internal/commons"
	"log/slog"
	"os"
)

// main is the entrypoint to the microservice. Here we:
// 1. Trigger dependency injection (to initialise everything that needs to be initialised)
// 2. Run the resulting gin engine to start the web server.
func main() {
	logger := commons.NewLogger(commons.LogFormatJSON, os.Stdout)
	logger.Info("Starting service...")

	var appConfig = &AppConfig{}
	err := loadAppConfigFromCommandFlags(appConfig, logger)
	if err != nil {
		logger.Error("Error occurred when parsing command line flags", slog.String("error", err.Error()))
		os.Exit(1)
	}

	engine, err := dependencyInjection(logger, appConfig)
	if err != nil {
		logger.Error("Error doing dependency injection to configure app", slog.String("error", err.Error()))
		os.Exit(2)
	}

	logger.Info("Starting application on port 8080")
	err = engine.Run(":8080")
	if err != nil {
		logger.Error("Error occurred when starting Gin app. Exiting", slog.String("error", err.Error()))
		os.Exit(3)
	}
}
