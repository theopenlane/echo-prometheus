package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	echo "github.com/theopenlane/echox"

	echoPrometheus "github.com/theopenlane/echo-prometheus"
)

func main() {
	e := echo.New()

	var configMetrics = echoPrometheus.NewConfig()
	configMetrics.Namespace = "namespace"
	configMetrics.NormalizeHTTPStatus = false
	configMetrics.Buckets = []float64{
		0.0005, // 0.5ms
		0.001,  // 1ms
		0.005,  // 5ms
		0.01,   // 10ms
		0.05,   // 50ms
		0.1,    // 100ms
		0.5,    // 500ms
		1,      // 1s
		2,      // 2s
	}

	e.Use(echoPrometheus.MetricsMiddlewareWithConfig(configMetrics))
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Error(e.Start(":1323"))
}
