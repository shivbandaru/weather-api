package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"weather-api/config"
	"weather-api/logger"
	"weather-api/models"
)

type usecase struct {
	cfg *config.MainConfig
}

// Weather Usecase constructor
func NewWeatherUsecase(cfg *config.MainConfig) *usecase {
	return &usecase{cfg: cfg}
}

func (u *usecase) GetWeather(ctx context.Context, lat, lon float64, appid string) (resp *models.WeatherData, err error) {

	// Construct the API URL reference https://openweathermap.org/current - API call section
	url := u.cfg.OpenWeather.Url + fmt.Sprintf("?lat=%.6f&lon=%.6f&appid=%s&units=imperial", lat, lon, appid)

	// Send HTTP GET request to the API
	response, err := http.Get(url)
	if err != nil {
		logger.Log.ErrorC(ctx, fmt.Sprintf("HTTP request failed with error: %v", err))
		return nil, err
	}

	if response.StatusCode != 200 {
		logger.Log.ErrorC(ctx, fmt.Sprintf("Non-OK response recieved from open weather api , status code : %d", response.StatusCode))
		return nil, fmt.Errorf(fmt.Sprintf("Non-OK:%d", response.StatusCode))
	}
	defer response.Body.Close()

	// Decode the JSON response
	var data map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		logger.Log.ErrorC(ctx, fmt.Sprintf("Failed to decode JSON: %v", err))
		return nil, err
	}

	// Extract weather information from the JSON data
	weatherDescription, temperature := extractWeatherInfo(data)
	visibility := extractVisibility(data)
	windSpeed := extractWindInfo(data)
	cloudCoverage := extractCloudCoverage(data)
	sunrise, sunset := extractSunriseSunset(data)

	// Classify weather type based on temperature
	weatherType := classifyWeather(temperature)

	//Log the response from the external api for quick reference
	logger.Log.InfoC(ctx, fmt.Sprintf("Response from Open Weather Api:%v", data))

	// Construct WeatherData struct and return
	return &models.WeatherData{
		WeatherDescription: weatherDescription,
		Temperature:        fmt.Sprintf("%v F", temperature),
		WeatherType:        weatherType,
		Visibility:         visibility,
		WindSpeed:          windSpeed,
		CloudCoverage:      cloudCoverage,
		Sunrise:            sunrise,
		Sunset:             sunset,
	}, nil
}

// Private helper methods
// extractWeatherInfo is a helper function that extracts weather description and temperature from the JSON data.
func extractWeatherInfo(data map[string]interface{}) (string, float64) {
	// Extract weather description from the 'weather' field
	weatherArray := data["weather"].([]interface{})
	weatherDescription := weatherArray[0].(map[string]interface{})["description"].(string)

	// Extract temperature from the 'main' field
	temperature := data["main"].(map[string]interface{})["temp"].(float64)

	return weatherDescription, temperature
}

// extractVisibility is a helper function that extracts visibility from the JSON data.
func extractVisibility(data map[string]interface{}) string {
	// Extract visibility from the 'visibility' field and convert to kilometers
	visibility := int(data["visibility"].(float64)) / 1000
	return fmt.Sprintf("%v Miles", visibility)
}

// extractWindInfo is a helper function that extracts wind speed and direction from the JSON data.
func extractWindInfo(data map[string]interface{}) string {
	// Extract wind speed from the 'wind' field
	windData := data["wind"].(map[string]interface{})
	windSpeed := windData["speed"].(float64)
	return fmt.Sprintf("%v miles/sec", windSpeed)
}

// extractCloudCoverage is a helper function that extracts cloud coverage from the JSON data.
func extractCloudCoverage(data map[string]interface{}) string {
	// Extract cloud coverage from the 'clouds' field
	cloudData := data["clouds"].(map[string]interface{})
	cloudCoverage := int(cloudData["all"].(float64))
	return fmt.Sprintf("%v%%", cloudCoverage)
}

// extractSunriseSunset is a helper function that extracts sunrise and sunset times from the JSON data.
func extractSunriseSunset(data map[string]interface{}) (time.Time, time.Time) {
	// Extract sunrise and sunset times from the 'sys' field
	sunriseUnix := int64(data["sys"].(map[string]interface{})["sunrise"].(float64))
	sunsetUnix := int64(data["sys"].(map[string]interface{})["sunset"].(float64))
	sunrise := time.Unix(sunriseUnix, 0)
	sunset := time.Unix(sunsetUnix, 0)
	return sunrise, sunset
}

// classifyWeather is a helper function that classifies the weather type based on temperature.
// TODO This classification can be set as a part of config in future to avoid making code changes
func classifyWeather(temperature float64) string {
	// Classify weather type based on temperature ranges
	if temperature <= 60 {
		return "cold"
	} else if temperature <= 85 {
		return "moderate"
	}
	return "hot"
}
