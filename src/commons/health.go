package commons

import "github.com/gin-gonic/gin"

func HealthController(ginEngine *gin.Engine) {
	ginEngine.GET("/management/health", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"status": "UP",
		})
	})
}
