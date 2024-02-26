package main

import (
	"log"

	"weather-api/config"
	"weather-api/server"
)

func main() {
	config := config.LoadConfig(".")

	server := server.NewServer(&config)

	err := server.Run()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}
