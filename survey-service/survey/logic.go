package survey

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator"
	"github.com/teris-io/shortid"
)

var (
	// ErrNotFound indicates that a requested survey was not found
	ErrNotFound = errors.New("Survey not found")

	// ErrInvalidRequest indicates that an invalid survey was provided
	ErrInvalidRequest = errors.New("Invalid survey input")
)

// ErrorResponse provides a structure for error responses
type ErrorResponse struct {
	Error string `json:"error"`
}

// surveyService implements the Service interface for managing surveys
type surveyService struct {
	repository Repository
	validator  *validator.Validate
}

// NewService creates a new survey service
// This function initializes a new surveyService with the given repository and a validator
func NewService(repository Repository) Service {
	return &surveyService{
		repository: repository,
		validator:  validator.New(),
	}
}

// Insert validates and inserts a new survey into the repository
// Generates a unique ID and sets the creation time for the survey
func (s *surveyService) Insert(survey *Survey) error {
	// Validate the survey structure
	if err := s.validator.Struct(survey); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidRequest, err)
	}

	// Generate a unique ID and set the creation time
	survey.ID = shortid.MustGenerate()
	survey.CreatedAt = time.Now().UTC().Unix()

	// Assign sequential IDs to survey questions
	for k := range survey.Questions {
		survey.Questions[k].ID = k + 1
	}

	// Insert the survey into the repository
	return s.repository.Insert(survey)
}

// LoadByID retrieves a survey by its ID from the repository
func (s *surveyService) LoadByID(id string) (*Survey, error) {
	return s.repository.LoadByID(id)
}

// Load retrieves all surveys from the repository
func (s *surveyService) Load() (*Surveys, error) {
	return s.repository.Load()
}
