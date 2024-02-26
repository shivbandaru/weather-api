package weather

import (
	"context"
	"weather-api/models"
)

type Usecase interface {
	GetWeather(ctx context.Context, lat, lon float64, appid string) (resp *models.WeatherData, err error)
}
