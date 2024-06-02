package repository

import (
	"errors"

	"github.com/VitaliySynytskyi/microservices-survey-app/survey-service/config"
	"github.com/VitaliySynytskyi/microservices-survey-app/survey-service/survey"
)

// NewSurveyRepository creates a new survey repository based on application configuration
// This function selects the appropriate repository type based on the configuration
func NewSurveyRepository(cfg config.Config) (survey.Repository, error) {
	switch cfg.Repository {
	case "memory":
		// Create a new in-memory repository
		return NewSurveyMemoryRepository()
	case "mongo":
		// Create a new MongoDB repository
		return NewSurveyMongoRepository(cfg.Mongo)
	default:
		// Return an error if no valid repository type is specified
		return nil, errors.New("no repository available")
	}
}
