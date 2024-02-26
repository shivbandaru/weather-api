package controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"weather-api/exception"
	"weather-api/logger"
	"weather-api/models"

	forecastUC "weather-api/forecast"

	"github.com/labstack/echo"
)

type apiHandler struct {
	forecastUC forecastUC.Usecase
}

// NewWeatherHandlers News handlers constructor
func NewWeatherHandlers(forecastUC forecastUC.Usecase) *apiHandler {
	return &apiHandler{forecastUC: forecastUC}
}

func (h *apiHandler) GetWeather(c echo.Context) error {

	ctx := context.WithValue(c.Request().Context(), models.RequestID, c.Response().Header().Get("RequestID"))

	var lat, lon float64
	var err error

	latString := c.QueryParam("lat")
	lonString := c.QueryParam("lon")

	if len(latString) == 0 || len(lonString) == 0 {
		logger.Log.ErrorC(ctx, fmt.Sprintf("Field validation error: lat: %s, lon: %s", latString, lonString))
		c.JSON(http.StatusBadRequest, exception.NewError(http.StatusBadRequest, "both lat, lon fields are mandatory"))
		return nil

	}

	// Parse latitude and longitude from the request URL query parameters
	lat, err = strconv.ParseFloat(latString, 64)
	if err != nil {
		logger.Log.ErrorC(ctx, fmt.Sprintf("latitude parsing error:%v", err.Error()))
		c.JSON(http.StatusBadRequest, exception.NewError(http.StatusBadRequest, "latitude parsing error, value should be decimal"))
		return nil
	}

	// Parse latitude and longitude from the request URL query parameters
	lon, err = strconv.ParseFloat(lonString, 64)
	if err != nil {
		logger.Log.ErrorC(ctx, fmt.Sprintf("longitude parsing error:%v", err.Error()))
		c.JSON(http.StatusBadRequest, exception.NewError(http.StatusBadRequest, "longitude parsing error, value should be decimal"))
		return nil
	}

	logger.Log.InfoC(ctx, fmt.Sprintf("Incoming Request: Latitude:%.6f, Longitude:%.6f", lat, lon))

	weatherInfo, err := h.forecastUC.GetWeather(ctx, lat, lon)
	if err != nil {
		c.JSON(http.StatusInternalServerError, exception.NewError(http.StatusInternalServerError, err.Error()))
		return nil
	}

	logger.Log.InfoC(ctx, fmt.Sprintf("Weather Response:%+v", weatherInfo))

	c.JSON(http.StatusOK, weatherInfo)
	return nil
}

func (h *apiHandler) HealthCheck(c echo.Context) error {
	ctx := context.WithValue(c.Request().Context(), models.RequestID, c.Response().Header().Get("RequestID"))
	logger.Log.InfoC(ctx, "Healthy")
	return c.JSON(http.StatusOK, "OK")
}
