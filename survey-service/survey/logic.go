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
	ErrInvalidRequest = errors.New("invalid survey input")

	// ErrSurveyExpired indicates that the survey has expired
	ErrSurveyExpired = errors.New("survey has expired")

	// ErrSurveyInactive indicates that the survey is not active
	ErrSurveyInactive = errors.New("survey is not active")

	// ErrQuestionConditionalLogic indicates an issue with conditional logic
	ErrQuestionConditionalLogic = errors.New("invalid conditional logic")
)

// ErrorResponse provides a structure for error responses
type ErrorResponse struct {
	Error string `json:"error"`
}

// SurveyResponse provides a structure for survey responses
type SurveyResponse struct {
	Survey SurveyWithStatus `json:"survey"`
}

// SurveysResponse provides a structure for multiple surveys response
type SurveysResponse struct {
	Surveys []SurveyWithStatus `json:"surveys"`
}

// surveyService implements the Service interface for managing surveys
type surveyService struct {
	repository Repository
	validator  *validator.Validate
}

// NewService creates a new survey service
// This function initializes a new surveyService with the given repository and a validator
func NewService(repository Repository) Service {
	s := &surveyService{
		repository: repository,
		validator:  validator.New(),
	}

	// Register custom validation for question types
	s.validator.RegisterValidation("questiontype", validateQuestionType)

	return s
}

// validateQuestionType validates that the question type is one of the supported types
func validateQuestionType(fl validator.FieldLevel) bool {
	qt := fl.Field().String()
	switch QuestionType(qt) {
	case QuestionTypeSingleChoice, QuestionTypeMultipleChoice, QuestionTypeText,
		QuestionTypeRating, QuestionTypeScale, QuestionTypeDate:
		return true
	default:
		return false
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

	// Set active flag to true by default if not specified
	if !survey.Active {
		survey.Active = true
	}

	// Validate each question based on its type
	for k := range survey.Questions {
		// Assign sequential IDs to survey questions
		survey.Questions[k].ID = k + 1

		// Validate based on question type
		if err := s.validateQuestionByType(&survey.Questions[k]); err != nil {
			return fmt.Errorf("%w: %v", ErrInvalidRequest, err)
		}

		// Validate conditional logic if present
		if survey.Questions[k].ConditionalLogic != nil {
			if err := s.validateConditionalLogic(survey.Questions[k].ConditionalLogic, k+1, survey.Questions); err != nil {
				return fmt.Errorf("%w: %v", ErrQuestionConditionalLogic, err)
			}
		}

		// Assign IDs to options if not already assigned
		if len(survey.Questions[k].Options) > 0 {
			for j := range survey.Questions[k].Options {
				survey.Questions[k].Options[j].ID = j + 1
			}
		}
	}

	// Insert the survey into the repository
	return s.repository.Insert(survey)
}

// validateQuestionByType validates a question based on its type
func (s *surveyService) validateQuestionByType(q *Question) error {
	switch q.Type {
	case QuestionTypeSingleChoice, QuestionTypeMultipleChoice:
		// Choice questions must have options
		if len(q.Options) < 2 {
			return errors.New("choice questions must have at least 2 options")
		}
	case QuestionTypeRating, QuestionTypeScale:
		// Scale questions must have min and max values
		if q.MinValue == nil || q.MaxValue == nil {
			return errors.New("rating/scale questions must have min and max values")
		}
		if *q.MinValue >= *q.MaxValue {
			return errors.New("min value must be less than max value")
		}
	}
	return nil
}

// validateConditionalLogic validates the conditional logic of a question
func (s *surveyService) validateConditionalLogic(logic *ConditionalLogic, currentQuestionID int, questions []Question) error {
	// Source question must exist and must come before this question
	if logic.SourceQuestionID >= currentQuestionID {
		return errors.New("conditional logic can only reference previous questions")
	}

	// Find the source question
	var sourceQuestion *Question
	for i := range questions {
		if questions[i].ID == logic.SourceQuestionID {
			sourceQuestion = &questions[i]
			break
		}
	}

	if sourceQuestion == nil {
		return errors.New("source question not found")
	}

	// If the logic references an option, check that it exists in the source question
	if logic.SourceOptionID > 0 {
		if sourceQuestion.Type != QuestionTypeSingleChoice && sourceQuestion.Type != QuestionTypeMultipleChoice {
			return errors.New("source option ID can only be used with choice questions")
		}

		optionExists := false
		for i := range sourceQuestion.Options {
			if sourceQuestion.Options[i].ID == logic.SourceOptionID {
				optionExists = true
				break
			}
		}

		if !optionExists {
			return errors.New("source option not found in source question")
		}
	}

	return nil
}

// LoadByID retrieves a survey by its ID from the repository
// Returns the survey with its current status
func (s *surveyService) LoadByID(id string) (*SurveyWithStatus, error) {
	survey, err := s.repository.LoadByID(id)
	if err != nil {
		return nil, err
	}

	return &SurveyWithStatus{
		Survey: survey,
		Status: survey.GetStatus(),
	}, nil
}

// Load retrieves all surveys from the repository
// Returns surveys with their current status
func (s *surveyService) Load() ([]SurveyWithStatus, error) {
	surveys, err := s.repository.Load()
	if err != nil {
		return nil, err
	}

	// Convert surveys to surveys with status
	surveysWithStatus := make([]SurveyWithStatus, 0, len(*surveys))
	for _, survey := range *surveys {
		surveysWithStatus = append(surveysWithStatus, SurveyWithStatus{
			Survey: survey,
			Status: survey.GetStatus(),
		})
	}

	return surveysWithStatus, nil
}

// LoadActive retrieves all active surveys that haven't expired
func (s *surveyService) LoadActive() ([]SurveyWithStatus, error) {
	surveys, err := s.repository.Load()
	if err != nil {
		return nil, err
	}

	// Filter active non-expired surveys
	activeNonExpired := make([]SurveyWithStatus, 0)
	for _, survey := range *surveys {
		if survey.Active && !survey.IsExpired() {
			activeNonExpired = append(activeNonExpired, SurveyWithStatus{
				Survey: survey,
				Status: SurveyStatusActive,
			})
		}
	}

	return activeNonExpired, nil
}

// Update updates an existing survey
func (s *surveyService) Update(id string, survey *Survey) error {
	// Ensure the ID is set correctly
	survey.ID = id

	// Validate the survey structure
	if err := s.validator.Struct(survey); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidRequest, err)
	}

	// Check if the survey exists
	existing, err := s.repository.LoadByID(id)
	if err != nil {
		return err
	}

	// Preserve creation time
	survey.CreatedAt = existing.CreatedAt

	// Validate questions
	for k := range survey.Questions {
		// Ensure question IDs are sequential
		survey.Questions[k].ID = k + 1

		// Validate based on question type
		if err := s.validateQuestionByType(&survey.Questions[k]); err != nil {
			return fmt.Errorf("%w: %v", ErrInvalidRequest, err)
		}

		// Validate conditional logic if present
		if survey.Questions[k].ConditionalLogic != nil {
			if err := s.validateConditionalLogic(survey.Questions[k].ConditionalLogic, k+1, survey.Questions); err != nil {
				return fmt.Errorf("%w: %v", ErrQuestionConditionalLogic, err)
			}
		}

		// Assign IDs to options if not already assigned
		if len(survey.Questions[k].Options) > 0 {
			for j := range survey.Questions[k].Options {
				survey.Questions[k].Options[j].ID = j + 1
			}
		}
	}

	// Update the survey in the repository
	return s.repository.Update(survey)
}

// ActivateSurvey activates a survey
func (s *surveyService) ActivateSurvey(id string) error {
	survey, err := s.repository.LoadByID(id)
	if err != nil {
		return err
	}

	survey.Active = true
	return s.repository.Update(survey)
}

// DeactivateSurvey deactivates a survey
func (s *surveyService) DeactivateSurvey(id string) error {
	survey, err := s.repository.LoadByID(id)
	if err != nil {
		return err
	}

	survey.Active = false
	return s.repository.Update(survey)
}

// SetExpirationDate sets the expiration date for a survey
func (s *surveyService) SetExpirationDate(id string, expiresAt int64) error {
	survey, err := s.repository.LoadByID(id)
	if err != nil {
		return err
	}

	// Validate that expiration date is in the future
	if expiresAt <= time.Now().UTC().Unix() {
		return fmt.Errorf("%w: expiration date must be in the future", ErrInvalidRequest)
	}

	survey.ExpiresAt = expiresAt
	return s.repository.Update(survey)
}

// DeleteSurvey deletes a survey by ID
func (s *surveyService) DeleteSurvey(id string) error {
	return s.repository.Delete(id)
}
