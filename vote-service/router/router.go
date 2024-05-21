package router

import (
	"github.com/VitaliySynytskyi/microservices-survey-app/vote-service/handler"
	"github.com/VitaliySynytskyi/microservices-survey-app/vote-service/middleware"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

// NewRouter creates a new router
// This function sets up the router with middleware and routes for the vote service
func NewRouter(h *handler.VoteHTTPHandler) *chi.Mux {
	r := chi.NewRouter()

	// Set up CORS middleware
	// This middleware allows cross-origin requests from any origin
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},                              // Allow all origins
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},         // Allow GET, POST, and OPTIONS methods
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}, // Allow specified headers
		ExposedHeaders:   []string{"Link"},                           // Expose specified headers
		AllowCredentials: true,                                       // Allow credentials
		MaxAge:           300,                                         // Max age for preflight requests
	}))

	// Set up vote routes
	// This route group handles all vote-related endpoints
	r.Route("/vote", func(r chi.Router) {
		r.Use(middleware.AddSerializer) // Add the serializer middleware
		r.Post("/", h.Vote)             // POST /vote - casts a vote
	})

	// Set up results routes
	// This route group handles all results-related endpoints
	r.Route("/results", func(r chi.Router) {
		r.Use(middleware.AddSerializer) // Add the serializer middleware
		r.Get("/{id}", h.GetResults)    // GET /results/{id} - retrieves results for a specific survey
	})

	return r
}
