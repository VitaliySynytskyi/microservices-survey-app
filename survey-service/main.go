package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/VitaliySynytskyi/microservices-survey-app/survey-service/config"
	"github.com/VitaliySynytskyi/microservices-survey-app/survey-service/handler"
	"github.com/VitaliySynytskyi/microservices-survey-app/survey-service/logger"
	"github.com/VitaliySynytskyi/microservices-survey-app/survey-service/repository"
	"github.com/VitaliySynytskyi/microservices-survey-app/survey-service/router"
	"github.com/VitaliySynytskyi/microservices-survey-app/survey-service/server"
	"github.com/VitaliySynytskyi/microservices-survey-app/survey-service/survey"
)

func main() {
	// Initialize the logger
	log := logger.NewConsoleLogger()

	// Load application configuration
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot load application configuration")
		os.Exit(1)
	}

	// Initialize the repository based on the configuration
	repo, err := repository.NewSurveyRepository(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot connect to repository")
		os.Exit(1)
	}

	// Initialize the survey service with the repository
	service := survey.NewService(repo)

	// Set up the HTTP server dependencies
	httpHandler := handler.NewSurveyHTTPHandler(service, &log)
	httpRouter := router.NewRouter(httpHandler)
	httpServer := server.NewHTTPServer(httpRouter, cfg.HTTP)

	// Start the HTTP server in a new goroutine
	go func() {
		log.Info().Str("on", httpServer.Addr).Msg("Starting HTTP server")
		err := httpServer.ListenAndServe()
		if err != nil {
			log.Fatal().Err(err).Msg("Server shutdown")
			os.Exit(1)
		}
	}()

	// Set up the gRPC server dependencies
	grpcHandler := handler.NewSurveyGrpcHandler(service, &log)

	// Start the gRPC server
	server.StartGrpcServer(grpcHandler, cfg.Grpc, &log)

	// Set up signal handling for graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	// Block until a signal is received
	sig := <-c
	log.Warn().Msgf("Signal received: %v", sig)

	// Gracefully shutdown the HTTP server, allowing up to 30 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	httpServer.Shutdown(ctx)
}
