package handler

import (
	"context"
	"errors"

	protos "github.com/VitaliySynytskyi/microservices-survey-app/survey-service/protos/survey"
	"github.com/VitaliySynytskyi/microservices-survey-app/survey-service/survey"
	"github.com/rs/zerolog"
)

// SurveyGrpcHandler handles gRPC requests for surveys
type SurveyGrpcHandler struct {
	service survey.Service
	log     *zerolog.Logger
}

// NewSurveyGrpcHandler creates a new survey gRPC handler
func NewSurveyGrpcHandler(service survey.Service, log *zerolog.Logger) *SurveyGrpcHandler {
	return &SurveyGrpcHandler{service, log}
}

// GetSurvey loads and returns a requested survey
func (g *SurveyGrpcHandler) GetSurvey(ctx context.Context, r *protos.SurveyRequest) (*protos.SurveyResponse, error) {
	id := r.GetId()
	g.log.Info().Str("id", id).Msg("GetSurvey request received")

	// Load survey by ID
	s, err := g.service.LoadByID(id)
	if err != nil {
		return handleSurveyError(g.log, id, err)
	}

	// Build survey response
	res := buildSurveyResponse(s)
	return res, nil
}

// handleSurveyError handles errors during survey loading
// Code Smell: Long Method
// Refactoring: Extract Method to handle error logging
func handleSurveyError(log *zerolog.Logger, id string, err error) (*protos.SurveyResponse, error) {
	if errors.Is(err, survey.ErrNotFound) {
		log.Debug().Str("id", id).Msg("Survey not found")
	} else {
		log.Error().Err(err).Str("id", id).Msg("Unable to load survey")
	}
	return nil, err
}

// buildSurveyResponse builds the response from the survey data
// Code Smell: Long Method
// Refactoring: Extract Method to build survey response
func buildSurveyResponse(s *survey.Survey) *protos.SurveyResponse {
	res := &protos.SurveyResponse{
		Id:        s.ID,
		Name:      s.Name,
		CreatedAt: s.CreatedAt,
	}
	for _, q := range s.Questions {
		res.Questions = append(res.Questions, &protos.QuestionResponse{
			Id:   int32(q.ID),
			Text: q.Text,
		})
	}
	return res
}
