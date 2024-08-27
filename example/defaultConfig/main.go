package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	echo "github.com/theopenlane/echox"

	echoPrometheus "github.com/theopenlane/echo-prometheus"
)

func main() {
	e := echo.New()

	e.Use(echoPrometheus.MetricsMiddleware())
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Error(e.Start(":1323"))
}
