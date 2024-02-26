package config

import (
	"weather-api/exception"

	"github.com/spf13/viper"
)

type MainConfig struct {
	Server      ServerConfig
	OpenWeather OpenWeatherConfig
}

type ServerConfig struct {
	Address string `mapstructure:"Address"`
}

type OpenWeatherConfig struct {
	Url string `mapstructure:"Url"`
}

func LoadConfig(path string) (config MainConfig) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	exception.PanicIfNeeded(err)

	err = viper.Unmarshal(&config)
	exception.PanicIfNeeded(err)

	return config
}
