package main

import (
	"github.com/Harmelodic/init-microservice-go/src/account"
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
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(sloggin.NewWithConfig(logger, sloggin.Config{
		WithTraceID: true,
	}))
	logger.Info("Gin engine configured")

	// Dependency Injection!
	dependencyInjection(engine)

	logger.Info("Starting application on port 8080")
	err := engine.Run(":8080")
	if err != nil {
		logger.Error("Error occurred when starting Gin app. Exiting")
		os.Exit(1)
	}
}

func dependencyInjection(engine *gin.Engine) {
	managementRoutes(engine)

	account.Controller(engine, account.Service{})
}
