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
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},                 // Allow all standard methods
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}, // Allow specified headers
		ExposedHeaders:   []string{"Link"},                                                    // Expose specified headers
		AllowCredentials: true,                                                                // Allow credentials
		MaxAge:           300,                                                                 // Max age for preflight requests
	}))

	// Survey routes
	// This route group handles all survey-related endpoints
	r.Route("/surveys", func(r chi.Router) {
		r.Use(middleware.AddSerializer) // Add the serializer middleware

		// Base survey endpoints
		r.Get("/", h.Collection)             // GET /surveys - retrieves all surveys
		r.Get("/active", h.ActiveCollection) // GET /surveys/active - retrieves all active surveys
		r.Post("/", h.Post)                  // POST /surveys - creates a new survey

		// Individual survey endpoints
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", h.Get)       // GET /surveys/{id} - retrieves a specific survey by ID
			r.Put("/", h.Put)       // PUT /surveys/{id} - updates a specific survey
			r.Delete("/", h.Delete) // DELETE /surveys/{id} - deletes a specific survey

			// Survey management endpoints
			r.Post("/activate", h.Activate)        // POST /surveys/{id}/activate - activates a survey
			r.Post("/deactivate", h.Deactivate)    // POST /surveys/{id}/deactivate - deactivates a survey
			r.Post("/expiration", h.SetExpiration) // POST /surveys/{id}/expiration - sets an expiration date
		})
	})

	return r
}
