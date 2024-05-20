package repository

import (
	"errors"

	"github.com/VitaliySynytskyi/microservices-survey-app/survey-service/config"
	"github.com/VitaliySynytskyi/microservices-survey-app/survey-service/survey"
)

// NewSurveyRepository creates a new survey repository based on application configuration
func NewSurveyRepository(cfg config.Config) (survey.Repository, error) {
	switch cfg.Repository {
	case "memory":
		return NewSurveyMemoryRepository()
	case "mongo":
		return NewSurveyMongoRepository(cfg.Mongo)
	}

	return nil, errors.New("No repository available")
}
