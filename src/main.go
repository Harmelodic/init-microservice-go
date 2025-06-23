package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	// TODO: Replace with real log
	fmt.Println("Starting service...")

	router := gin.Default()

	// TODO: Move this out to a management package for health checks, observability scraping and mgmt endpoints.
	router.GET("/management/health", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"status": "UP",
		})
	})

	err := router.Run()
	if err != nil {
		// TODO: Replace with real log
		fmt.Println("Error occurred when starting Gin app. Exiting")
		os.Exit(1)
	}
}
