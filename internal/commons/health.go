package commons

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthIndicator describes how a health indicator should function. A component can be a HealthIndicator if it has a
// name (for visibility) and can report if it is healthy or not.
type HealthIndicator interface {
	IndicateHealth() (name string, isHealthy bool)
}

// healthReport is the API response for reporting what is HEALTHY and FAILING.
type healthReport struct {
	Indicator string `json:"indicator"`
	Status    Status `json:"status"`
}

// Status is an enumeration for conveying the health status.
type Status string

const (
	// HEALTHY conveys a Status that is up.
	HEALTHY Status = "HEALTHY"
	// FAILING conveys a Status that is down.
	FAILING Status = "FAILING"
)

// LivenessController creates an endpoint to report whether the service is HEALTHY (HTTP 200) or FAILING (HTTP 503) in
// regard to
// "liveness": https://kubernetes.io/docs/concepts/configuration/liveness-readiness-startup-probes/#liveness-probe.
//
// To report this, it uses a default healthReport, plus any HealthIndicator structs provided to it.
// If you build something that you want to include in the LivenessController, first ensure it can be a HealthIndicator,
// then simply pass it in when doing dependency injection.
func LivenessController(ginEngine *gin.Engine, indicators ...HealthIndicator) {
	ginEngine.GET("/health/liveness", func(context *gin.Context) {
		var shouldGenerallyReportUp = true

		healthChecks := make([]healthReport, len(indicators)+1)
		healthChecks[0] = healthReport{
			Indicator: "LivenessController",
			Status:    HEALTHY,
		}

		for index, indicator := range indicators {
			name, isHealthy := indicator.IndicateHealth()
			newCheck := healthReport{
				Indicator: name,
				Status:    boolToStatus(isHealthy),
			}

			if newCheck.Status == FAILING {
				shouldGenerallyReportUp = false
			}

			healthChecks[index+1] = newCheck
		}

		if shouldGenerallyReportUp {
			context.JSON(http.StatusOK, healthChecks)
		} else {
			context.JSON(http.StatusServiceUnavailable, healthChecks)
		}
	})
}

// ReadinessController creates an endpoint to report whether the service is HEALTHY (HTTP 200) or FAILING (HTTP 503) in
// regard to
// "readiness": https://kubernetes.io/docs/concepts/configuration/liveness-readiness-startup-probes/#readiness-probe.
//
// To report this, it uses a default healthReport, plus any HealthIndicator structs provided to it.
// If you build something that you want to include in the LivenessController, first ensure it can be a HealthIndicator,
// then simply pass it in when doing dependency injection.
func ReadinessController(ginEngine *gin.Engine, indicators ...HealthIndicator) {
	ginEngine.GET("/health/readiness", func(context *gin.Context) {
		var shouldReportUp = true

		healthChecks := make([]healthReport, len(indicators)+1)
		healthChecks[0] = healthReport{
			Indicator: "ReadinessController",
			Status:    HEALTHY,
		}

		for index, indicator := range indicators {
			name, isHealthy := indicator.IndicateHealth()
			newCheck := healthReport{
				Indicator: name,
				Status:    boolToStatus(isHealthy),
			}

			if newCheck.Status == FAILING {
				shouldReportUp = false
			}

			healthChecks[index+1] = newCheck
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
		return HEALTHY
	}

	return FAILING
}
