package router

import (
	"github.com/VitaliySynytskyi/microservices-survey-app/survey-service/handler"
	"github.com/VitaliySynytskyi/microservices-survey-app/survey-service/middleware"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

// NewRouter creates a new router
// This function initializes the router and sets up the routes and middleware
func NewRouter(h handler.SurveyHTTPHandler) *chi.Mux {
	r := chi.NewRouter()

	// Set up CORS middleware
	// This middleware allows cross-origin requests from any origin
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},                                                       // Allow all origins
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},                                  // Allow GET, POST, and OPTIONS methods
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}, // Allow specified headers
		ExposedHeaders:   []string{"Link"},                                                    // Expose specified headers
		AllowCredentials: true,                                                                // Allow credentials
		MaxAge:           300,                                                                 // Max age for preflight requests
	}))

	// Survey routes
	// This route group handles all survey-related endpoints
	r.Route("/surveys", func(r chi.Router) {
		r.Use(middleware.AddSerializer) // Add the serializer middleware
		r.Get("/", h.Collection)        // GET /surveys - retrieves all surveys
		r.Get("/{id}", h.Get)           // GET /surveys/{id} - retrieves a specific survey by ID
		r.Post("/", h.Post)             // POST /surveys - creates a new survey
	})

	return r
}
