package controller

import (
	"github.com/labstack/echo"
)

func (h *apiHandler) RegisterHTTPEndpoints(e *echo.Echo) {

	e.GET("/healthcheck", h.HealthCheck)

	group := e.Group("/api/v1/weather")
	group.GET("/forecast", h.GetWeather)

}
