package server

import (
	"fmt"
	"net"
	"os"

	"github.com/VitaliySynytskyi/microservices-survey-app/survey-service/config"
	"github.com/VitaliySynytskyi/microservices-survey-app/survey-service/handler"
	protos "github.com/VitaliySynytskyi/microservices-survey-app/survey-service/protos/survey"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// StartGrpcServer starts a new gRPC server
// This function sets up and starts a gRPC server for the Survey service
func StartGrpcServer(h *handler.SurveyGrpcHandler, cfg config.GrpcConfig, log *zerolog.Logger) {
	// Create a new gRPC server
	gs := grpc.NewServer()

	// Register the Survey gRPC handler with the server
	protos.RegisterSurveyServer(gs, h)

	// Register reflection service on gRPC server
	reflection.Register(gs)

	// Construct the address from the hostname and port in the configuration
	addr := fmt.Sprintf("%s:%d", cfg.Hostname, cfg.Port)

	// Listen on the specified network address
	l, err := net.Listen(cfg.Network, addr)
	if err != nil {
		// Log the error and exit if the server cannot start
		log.Fatal().Err(err).Msg("Unable to start gRPC server.")
		os.Exit(1)
	}

	// Log the start of the gRPC server
	log.Info().Str("on", addr).Msg("Starting gRPC server")

	// Serve gRPC requests on the listener
	gs.Serve(l)
}
