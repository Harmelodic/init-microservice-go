package commons

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HealthIndicator describes how a health indicator should function. A component can be a HealthIndicator if it has a
// name (for visibility) and can report if it is healthy or not
type HealthIndicator interface {
	Name() string
	IsHealthy() bool
}

type Status string

const (
	UP   Status = "UP"
	DOWN Status = "DOWN"
)

// healthReport is the API response for reporting which indicators are UP and DOWN
type healthReport struct {
	Indicator string `json:"indicator"`
	Status    Status `json:"status"`
}

// LivenessController creates an endpoint to report whether the service is UP (HTTP 200) or DOWN (HTTP 503) in regard to
// "liveness": https://kubernetes.io/docs/concepts/configuration/liveness-readiness-startup-probes/#liveness-probe.
// To report this, it uses a default healthReport, plus any HealthIndicator structs provided to it.
// If you build something that you want to include in the LivenessController, first ensure it can be a HealthIndicator,
// then simply pass it in when doing dependency injection.
func LivenessController(ginEngine *gin.Engine, indicators ...HealthIndicator) {
	ginEngine.GET("/health/liveness", func(context *gin.Context) {
		var shouldReportUp = true

		healthChecks := make([]healthReport, len(indicators)+1)
		healthChecks[0] = healthReport{
			Indicator: "LivenessController",
			Status:    UP,
		}

		for i, indicator := range indicators {
			newCheck := healthReport{
				Indicator: indicator.Name(),
				Status:    boolToStatus(indicator.IsHealthy()),
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

// ReadinessController creates an endpoint to report whether the service is UP (HTTP 200) or DOWN (HTTP 503) in regard
// to "readiness": https://kubernetes.io/docs/concepts/configuration/liveness-readiness-startup-probes/#readiness-probe.
// To report this, it uses a default healthReport, plus any HealthIndicator structs provided to it.
// If you build something that you want to include in the LivenessController, first ensure it can be a HealthIndicator,
// then simply pass it in when doing dependency injection.
func ReadinessController(ginEngine *gin.Engine, indicators ...HealthIndicator) {
	ginEngine.GET("/health/readiness", func(context *gin.Context) {
		var shouldReportUp = true

		healthChecks := make([]healthReport, len(indicators)+1)
		healthChecks[0] = healthReport{
			Indicator: "ReadinessController",
			Status:    UP,
		}

		for i, indicator := range indicators {
			newCheck := healthReport{
				Indicator: indicator.Name(),
				Status:    boolToStatus(indicator.IsHealthy()),
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

func boolToStatus(up bool) Status {
	if up {
		return UP
	}
	return DOWN
}
