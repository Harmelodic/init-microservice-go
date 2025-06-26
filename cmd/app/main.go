package main

import (
	"github.com/Harmelodic/init-microservice-go/internal/account"
	"github.com/Harmelodic/init-microservice-go/internal/commons"
	"github.com/gin-gonic/gin"
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

	dependencyInjection(engine)

	logger.Info("Starting application on port 8080")
	err := engine.Run(":8080")
	if err != nil {
		logger.Error("Error occurred when starting Gin app. Exiting")
		os.Exit(1)
	}
}

func dependencyInjection(engine *gin.Engine) {
	accountRepository := account.DefaultRepository{}
	accountService := account.DefaultService{Repository: &accountRepository}

	account.Controller(engine, &accountService)

	commons.LivenessController(engine)
	commons.ReadinessController(engine, &accountRepository)
}
