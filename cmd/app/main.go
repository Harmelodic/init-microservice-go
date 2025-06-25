package main

import (
	"github.com/Harmelodic/init-microservice-go/internal/account"
	"github.com/Harmelodic/init-microservice-go/internal/commons"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	"os"
)

// main is the entrypoint to the microservice. Here we:
// 1. Configure the gin engine
// 2. Trigger dependency injection
// 3. Run the gin engine to start the web server.
func main() {
	logger := commons.NewLogger()
	logger.Info("Starting service...")

	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(sloggin.NewWithConfig(logger, sloggin.Config{
		WithTraceID: true,
	}))
	logger.Info("Gin engine configured")

	dependencyInjection(engine)

	logger.Info("Starting application on port 8080")
	err := engine.Run(":8080")
	if err != nil {
		logger.Error("Error occurred when starting Gin app. Exiting")
		os.Exit(1)
	}
}

func dependencyInjection(engine *gin.Engine) {
	account.Controller(engine, &account.Service{})
	commons.HealthController(engine)
}
