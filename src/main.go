package main

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"os"
)

func main() {
	slog.Info("Starting service...")

	router := gin.Default()

	// TODO: Move this out to a management package for health checks, observability scraping and mgmt endpoints.
	router.GET("/management/health", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"status": "UP",
		})
	})

	err := router.Run()
	if err != nil {
		slog.Error("Error occurred when starting Gin app. Exiting")
		os.Exit(1)
	}
}
