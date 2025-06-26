package commons

import (
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	"log/slog"
)

func NewGinEngine(logger *slog.Logger) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(sloggin.NewWithConfig(logger, sloggin.Config{
		WithTraceID: true,
	}))
	logger.Info("Gin engine configured")
	return engine
}
