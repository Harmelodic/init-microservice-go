package commons

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthIndicator interface {
	Name() string
	CheckHealth() Status
}

type Status string

const (
	UP   Status = "UP"
	DOWN Status = "DOWN"
)

// healthCheck is the API response for reporting which indicators are UP and DOWN
type healthCheck struct {
	Indicator string `json:"indicator"`
	Status    Status `json:"status"`
}

// HealthController handles reporting whether the service should report UP (HTTP 200) or DOWN (HTTP 503).
// To report this, it uses a default healthCheck, plus any HealthIndicator structs provided to it.
// If you build something that you want to include in the HealthController, first ensure it can be a HealthIndicator,
// then simply pass it in when doing dependency injection.
func HealthController(ginEngine *gin.Engine, indicators ...HealthIndicator) {
	ginEngine.GET("/management/health", func(context *gin.Context) {
		var shouldReportUp = true

		healthChecks := make([]healthCheck, len(indicators)+1)
		healthChecks[0] = healthCheck{
			Indicator: "HealthController",
			Status:    UP,
		}

		for i, indicator := range indicators {
			newCheck := healthCheck{
				Indicator: indicator.Name(),
				Status:    indicator.CheckHealth(),
			}

			if newCheck.Status == DOWN {
				shouldReportUp = false
			}

			healthChecks[i+1] = newCheck
		}

		if shouldReportUp {
			context.JSON(http.StatusOK, healthChecks)
		} else {
			context.JSON(http.StatusServiceUnavailable, healthChecks)
		}
	})
}
