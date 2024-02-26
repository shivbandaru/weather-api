package weather

import (
	"context"
	"weather-api/models"
)

type Usecase interface {
	GetWeather(ctx context.Context, lat, lon float64) (resp *models.WeatherData, err error)
}
