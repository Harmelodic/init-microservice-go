package main

import (
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	"os"
)

// main entrypoint for booting the service, including:
// - Dependency injection
// - HTTP server listening
func main() {
	logger := getLogger()
	logger.Info("Starting service...")

	gin.SetMode(gin.ReleaseMode)
	ginEngine := gin.New()
	ginEngine.Use(gin.Recovery())
	ginEngine.Use(sloggin.NewWithConfig(logger, sloggin.Config{
		WithTraceID: true,
	}))
	logger.Info("Gin engine configured")

	registerManagementRoutes(ginEngine)
	logger.Info("Endpoints registered")

	logger.Info("Starting application on port 8080")
	err := ginEngine.Run(":8080")
	if err != nil {
		logger.Error("Error occurred when starting Gin app. Exiting")
		os.Exit(1)
	}
}
