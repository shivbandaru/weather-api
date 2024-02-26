package models

import "time"

// WeatherData represents the structure of weather data obtained from the OpenWeatherMap API.
// It is constructed based on the JSON response format documented at https://openweathermap.org/current.
type WeatherData struct {
	WeatherDescription string    `json:"weather_condition"` // Description of the weather condition
	Temperature        string    `json:"temperature"`       // Temperature in Celsius
	WeatherType        string    `json:"weather_type"`      // Type of weather condition (e.g., cold, moderate, hot)
	Visibility         string    `json:"visibility"`        // Visibility in kilometers
	WindSpeed          string    `json:"wind_speed"`        // Wind speed in meters per second
	CloudCoverage      string    `json:"cloud_coverage"`    // Cloud coverage in percentage
	Sunrise            time.Time `json:"sunrise"`           // Time of sunrise
	Sunset             time.Time `json:"sunset"`            // Time of sunset
}
