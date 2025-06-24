package account

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Controller configures appropriate HTTP routes
func Controller(engine *gin.Engine, service Service) {
	engine.GET("/v1/account", func(context *gin.Context) {
		accounts := service.GetAllAccounts()
		context.JSON(http.StatusOK, accounts)
	})
}
