package account

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

// Controller configures appropriate HTTP routes.
func Controller(engine *gin.Engine, service Service, logger *slog.Logger) {
	engine.GET("/v1/account", func(context *gin.Context) {
		accounts, err := service.GetAllAccounts()
		if err != nil {
			logger.Error(err.Error())
			context.Status(http.StatusInternalServerError)

			return
		}

		context.JSON(http.StatusOK, accounts)
	})
}
