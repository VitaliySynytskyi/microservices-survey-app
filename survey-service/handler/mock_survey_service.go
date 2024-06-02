package handler

import (
	"github.com/VitaliySynytskyi/microservices-survey-app/survey-service/survey"
	"github.com/stretchr/testify/mock"
)

// MockSurveyService is a mock implementation of the survey.Service interface
type MockSurveyService struct {
	mock.Mock
}

func (m *MockSurveyService) LoadByID(id string) (*survey.Survey, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*survey.Survey), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockSurveyService) Insert(s *survey.Survey) error {
	args := m.Called(s)
	return args.Error(0)
}

func (m *MockSurveyService) Update(s *survey.Survey) error {
	args := m.Called(s)
	return args.Error(0)
}

func (m *MockSurveyService) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockSurveyService) Load() (*survey.Surveys, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).(*survey.Surveys), args.Error(1)
	}
	return nil, args.Error(1)
}

// MockSerializer is a mock implementation of the survey.Serializer interface
type MockSerializer struct {
	mock.Mock
}

func (m *MockSerializer) Encode(s *survey.Survey) ([]byte, error) {
	args := m.Called(s)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockSerializer) EncodeMultiple(surveys *survey.Surveys) ([]byte, error) {
	args := m.Called(surveys)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockSerializer) Decode(data []byte) (*survey.Survey, error) {
	args := m.Called(data)
	return args.Get(0).(*survey.Survey), args.Error(1)
}

func (m *MockSerializer) DecodeMultiple(data []byte) (*survey.Surveys, error) {
	args := m.Called(data)
	return args.Get(0).(*survey.Surveys), args.Error(1)
}

func (m *MockSerializer) EncodeErrorResponse(err survey.ErrorResponse) ([]byte, error) {
	args := m.Called(err)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockSerializer) GetContentType() string {
	return "application/json"
}
