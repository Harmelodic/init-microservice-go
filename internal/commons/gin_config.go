package commons

import (
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"log/slog"
)

// NewGinEngine is a factory method for creating a production-ready gin.Engine instance.
func NewGinEngine(serviceName string, logger *slog.Logger) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(otelgin.Middleware(serviceName))
	engine.Use(sloggin.NewWithConfig(logger, sloggin.Config{
		WithTraceID: true,
	}))

	return engine
}
