package vote

import (
	"context"
	"errors"
	"fmt"
	"time"

	protos "github.com/VitaliySynytskyi/microservices-survey-app/survey-service/protos/survey"
	"github.com/go-playground/validator"
	uuid "github.com/satori/go.uuid"
)

var (
	// ErrInvalidRequest indicates that an invalid vote was provided
	ErrInvalidRequest = errors.New("Invalid vote input")

	// ErrResultsNotFound indicates that results could not be found for a given survey
	ErrResultsNotFound = errors.New("Results not found")
)

// ErrorResponse provides a structure for error responses
type ErrorResponse struct {
	Error string `json:"error"`
}

// voteService implements the Service interface for managing votes
type voteService struct {
	writer    WriterRepository
	results   ResultsRepository
	validator *validator.Validate
	surveys   protos.SurveyClient
}

// NewService creates a new vote service
// This function initializes and returns a new voteService instance
func NewService(w WriterRepository, r ResultsRepository, cli protos.SurveyClient) Service {
	return &voteService{
		writer:    w,
		results:   r,
		validator: validator.New(),
		surveys:   cli,
	}
}

// Insert validates and inserts a new vote
// This method fetches the survey to validate the vote, generates an ID, and stores the vote
func (s *voteService) Insert(v *Vote) error {
	// Validate the vote structure
	if err := s.validator.Struct(v); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidRequest, err)
	}

	// Fetch and validate the survey and question ID
	if err := s.validateSurveyAndQuestion(v); err != nil {
		return err
	}

	// Generate a unique ID and set the timestamp
	v.ID = uuid.NewV4().String()
	v.Timestamp = time.Now().UTC().Unix()

	// Insert the vote into the repository
	return s.writer.Insert(v)
}

// validateSurveyAndQuestion validates the survey and question ID
// This method fetches the survey and checks if the question ID is valid
func (s *voteService) validateSurveyAndQuestion(v *Vote) error {
	// Fetch the survey
	req := &protos.SurveyRequest{Id: v.Survey}
	surv, err := s.surveys.GetSurvey(context.Background(), req)
	if err != nil {
		return ErrInvalidRequest
	}

	// Validate the question ID
	if !s.isValidQuestionID(v.Question, surv.GetQuestions()) {
		return ErrInvalidRequest
	}

	return nil
}

// isValidQuestionID checks if the question ID is valid within the survey questions
// This method returns true if the question ID is found in the survey questions
func (s *voteService) isValidQuestionID(questionID int, questions []*protos.QuestionResponse) bool {
	for _, q := range questions {
		if q.GetId() == int32(questionID) {
			return true
		}
	}
	return false
}

// GetResults retrieves the results for a given survey ID
// This method fetches the results from the results repository
func (s *voteService) GetResults(surveyID string) (Results, error) {
	return s.results.GetResults(surveyID)
}
