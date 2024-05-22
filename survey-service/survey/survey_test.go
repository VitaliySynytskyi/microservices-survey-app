package survey

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRepository is a mock implementation of the Repository interface
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Insert(survey *Survey) error {
	args := m.Called(survey)
	return args.Error(0)
}

func (m *MockRepository) LoadByID(id string) (*Survey, error) {
	args := m.Called(id)
	return args.Get(0).(*Survey), args.Error(1)
}

func (m *MockRepository) Load() (*Surveys, error) {
	args := m.Called()
	return args.Get(0).(*Surveys), args.Error(1)
}

// TestNewService tests the creation of a new survey service
func TestNewService(t *testing.T) {
	repo := new(MockRepository)
	service := NewService(repo)

	assert.NotNil(t, service, "Expected new service to be non-nil")
	assert.IsType(t, &surveyService{}, service, "Expected service to be of type surveyService")
}

// TestInsert tests the insertion of a new survey
func TestInsert(t *testing.T) {
	repo := new(MockRepository)
	service := NewService(repo)

	// Create a valid survey
	validSurvey := &Survey{
		Name: "Test Survey",
		Questions: []Question{
			{Text: "What is your name?"},
			{Text: "How old are you?"},
		},
	}

	// Mock the repository response
	repo.On("Insert", validSurvey).Return(nil)

	// Insert the survey
	err := service.Insert(validSurvey)

	// Validate the results
	assert.NoError(t, err, "Expected no error on valid survey insert")
	assert.NotEmpty(t, validSurvey.ID, "Expected survey ID to be set")
	assert.NotZero(t, validSurvey.CreatedAt, "Expected survey creation time to be set")
	assert.Equal(t, 1, validSurvey.Questions[0].ID, "Expected question ID to be set")
	assert.Equal(t, 2, validSurvey.Questions[1].ID, "Expected question ID to be set")
	repo.AssertExpectations(t)
}

// TestInsertInvalid tests the insertion of an invalid survey
func TestInsertInvalid(t *testing.T) {
	repo := new(MockRepository)
	service := NewService(repo)

	// Create an invalid survey (missing name)
	invalidSurvey := &Survey{
		Questions: []Question{
			{Text: "What is your name?"},
			{Text: "How old are you?"},
		},
	}

	// Insert the survey
	err := service.Insert(invalidSurvey)

	// Validate the results
	assert.Error(t, err, "Expected error on invalid survey insert")
	assert.True(t, errors.Is(err, ErrInvalidRequest), "Expected error to be ErrInvalidRequest")
	repo.AssertNotCalled(t, "Insert", invalidSurvey)
}

// TestLoadByID tests loading a survey by ID
func TestLoadByID(t *testing.T) {
	repo := new(MockRepository)
	service := NewService(repo)

	// Create a mock survey
	mockSurvey := &Survey{
		ID:        "testID",
		Name:      "Test Survey",
		CreatedAt: time.Now().UTC().Unix(),
		Questions: []Question{
			{ID: 1, Text: "What is your name?"},
			{ID: 2, Text: "How old are you?"},
		},
	}

	// Mock the repository response
	repo.On("LoadByID", "testID").Return(mockSurvey, nil)

	// Load the survey
	survey, err := service.LoadByID("testID")

	// Validate the results
	assert.NoError(t, err, "Expected no error on load by ID")
	assert.Equal(t, mockSurvey, survey, "Expected loaded survey to match mock survey")
	repo.AssertExpectations(t)
}

// TestLoad tests loading all surveys
func TestLoad(t *testing.T) {
	repo := new(MockRepository)
	service := NewService(repo)

	// Create mock surveys
	mockSurveys := &Surveys{
		&Survey{
			ID:        "testID1",
			Name:      "Test Survey 1",
			CreatedAt: time.Now().UTC().Unix(),
			Questions: []Question{
				{ID: 1, Text: "What is your name?"},
				{ID: 2, Text: "How old are you?"},
			},
		},
		&Survey{
			ID:        "testID2",
			Name:      "Test Survey 2",
			CreatedAt: time.Now().UTC().Unix(),
			Questions: []Question{
				{ID: 1, Text: "What is your favorite color?"},
				{ID: 2, Text: "What is your favorite food?"},
			},
		},
	}

	// Mock the repository response
	repo.On("Load").Return(mockSurveys, nil)

	// Load the surveys
	surveys, err := service.Load()

	// Validate the results
	assert.NoError(t, err, "Expected no error on load")
	assert.Equal(t, mockSurveys, surveys, "Expected loaded surveys to match mock surveys")
	repo.AssertExpectations(t)
}
