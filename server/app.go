package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo"

	"weather-api/config"
	"weather-api/exception"
	"weather-api/forecast/controller"
	"weather-api/forecast/usecase"
	"weather-api/middleware"
)

type Server struct {
	//httpServer *http.Server
	cfg *config.MainConfig
}

func NewServer(cfg *config.MainConfig) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Run() error {

	//Create a signal channel to capture interrupt and terminate signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	//Crete a new echo
	e := echo.New()
	midl := middleware.InitMiddleware()
	e.Use(midl.RequestID)

	weatherUC := usecase.NewWeatherUsecase(s.cfg)
	handler := controller.NewWeatherHandlers(weatherUC)
	handler.RegisterHTTPEndpoints(e)

	go func() {
		if err := e.Start(s.cfg.Server.Address); err != nil && err != http.ErrServerClosed {
			exception.PanicIfNeeded(err)
		}
	}()

	fmt.Printf("Server started on %s", s.cfg.Server.Address)

	//wait for interrupt or termination signal
	<-signalChan

	//Create a deadline for shutting down the server gracefully
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Shutdown the server gracefully
	if err := e.Shutdown(ctx); err != nil {
		fmt.Println("Failed to gracefully shutdown the server:", err)
		os.Exit(1)
	}

	fmt.Println("server gracefully shutdown")

	return nil
}
