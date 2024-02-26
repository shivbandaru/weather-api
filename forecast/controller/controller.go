package controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
	appid := c.QueryParam("appid")

	if len(latString) == 0 || len(lonString) == 0 {
		logger.Log.ErrorC(ctx, fmt.Sprintf("Field validation error: lat: %s, lon: %s", latString, lonString))
		c.JSON(http.StatusBadRequest, exception.NewError(http.StatusBadRequest, "both lat, lon fields are mandatory"))
		return nil

	}

	//apikey validation
	if len(appid) == 0 {
		logger.Log.ErrorC(ctx, fmt.Sprintf("Field validation error: appid is mandatory, appid: %s", appid))
		c.JSON(http.StatusBadRequest, exception.NewError(http.StatusBadRequest, "appid is mandatory"))
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

	weatherInfo, err := h.forecastUC.GetWeather(ctx, lat, lon, appid)

	//Below Code check the status code on the open api response and act accordingly , this needs to be revisited
	if err != nil {
		if strings.Contains(err.Error(), "Non-OK:") {
			switch strings.SplitAfter(err.Error(), ":")[1] {
			case "401":
				c.JSON(http.StatusUnauthorized, exception.NewError(http.StatusUnauthorized, "Unauthorized"))
			case "400":
				c.JSON(http.StatusBadRequest, exception.NewError(http.StatusBadRequest, "Invalid Input"))
			case "429":
				c.JSON(http.StatusTooManyRequests, exception.NewError(http.StatusTooManyRequests, "Too Many Requests"))
			default:
				c.JSON(http.StatusInternalServerError, exception.NewError(http.StatusInternalServerError, "Unknown Error"))
			}

		} else {
			c.JSON(http.StatusInternalServerError, exception.NewError(http.StatusInternalServerError, err.Error()))
		}
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
