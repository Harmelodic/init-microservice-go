package main

import "github.com/gin-gonic/gin"

func managementRoutes(ginEngine *gin.Engine) {
	ginEngine.GET("/management/health", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"status": "UP",
		})
	})
}
