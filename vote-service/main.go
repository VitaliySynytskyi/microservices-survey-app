package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	protos "github.com/VitaliySynytskyi/microservices-survey-app/survey-service/protos/survey"
	"google.golang.org/grpc"

	"github.com/VitaliySynytskyi/microservices-survey-app/vote-service/config"
	"github.com/VitaliySynytskyi/microservices-survey-app/vote-service/handler"
	"github.com/VitaliySynytskyi/microservices-survey-app/vote-service/logger"
	"github.com/VitaliySynytskyi/microservices-survey-app/vote-service/repository"
	"github.com/VitaliySynytskyi/microservices-survey-app/vote-service/router"
	"github.com/VitaliySynytskyi/microservices-survey-app/vote-service/serializer"
	"github.com/VitaliySynytskyi/microservices-survey-app/vote-service/server"
	"github.com/VitaliySynytskyi/microservices-survey-app/vote-service/vote"
)

func main() {
	// Get a logger
	log := logger.NewConsoleLogger()

	// Load application configuration
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot load application configuration")
		os.Exit(1)
	}

	// Load the write repository
	sz := serializer.NewVoteJSONSerializer()
	writer, err := repository.NewRabbitVoteWriterRepository(cfg.Rabbit, sz)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot connect to write repository")
		os.Exit(1)
	}

	// Load the results repository
	results, err := repository.NewPostgresResultsRepository(cfg.Postgres)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot connect to results repository")
		os.Exit(1)
	}

	// Connect to the survey gRPC service
	grpcAddr := fmt.Sprintf("%s:%d", cfg.SurveyGrpc.Hostname, cfg.SurveyGrpc.Port)
	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot connect to survey gRPC service")
		os.Exit(1)
	}
	cli := protos.NewSurveyClient(conn)
	log.Info().Str("on", grpcAddr).Msg("Connected to survey gRPC service")

	// Load the service
	service := vote.NewService(writer, results, cli)

	// Load HTTP dependencies
	httpHandler := handler.NewVoteHTTPHandler(service, &log)
	httpRouter := router.NewRouter(httpHandler)
	httpServer := server.NewHTTPServer(httpRouter, cfg.HTTP)

	// Start the HTTP server
	go func() {
		log.Info().Str("on", httpServer.Addr).Msg("Starting HTTP server")
		err := httpServer.ListenAndServe()
		if err != nil {
			log.Fatal().Err(err).Msg("Server shutdown")
			os.Exit(1)
		}
	}()

	// Listen for sigterm or interupt signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Block until a signal is received
	sig := <-c
	log.Warn().Msgf("Signal received: %v", sig)

	// Gracefully shutdown the server allowing up to 30 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	httpServer.Shutdown(ctx)
}
