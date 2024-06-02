package server

import (
	"fmt"
	"net/http"

	"github.com/VitaliySynytskyi/microservices-survey-app/vote-service/config"
	"github.com/go-chi/chi"
)

// NewHTTPServer creates a new HTTP server
// This function sets up and returns an HTTP server configured with the specified router and settings
func NewHTTPServer(router *chi.Mux, cfg config.HTTPConfig) *http.Server {
	// Create a new HTTP server with the specified configuration
	return &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.Hostname, cfg.Port), // Server address (hostname:port)
		Handler:      router,                                      // Router handling incoming requests
		ReadTimeout:  cfg.ReadTimeout,                             // Maximum duration for reading the request
		WriteTimeout: cfg.WriteTimeout,                            // Maximum duration for writing the response
		IdleTimeout:  cfg.IdleTimeout,                             // Maximum amount of time to wait for the next request
	}
}
