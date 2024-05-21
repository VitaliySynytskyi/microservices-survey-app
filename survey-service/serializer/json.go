package serializer

import (
	"encoding/json"

	"github.com/VitaliySynytskyi/microservices-survey-app/survey-service/survey"
)

// surveyJSONSerializer is an implementation of the survey.Serializer interface for JSON
type surveyJSONSerializer struct{}

// NewSurveyJSONSerializer creates a new survey JSON serializer
// This function returns a new instance of surveyJSONSerializer
func NewSurveyJSONSerializer() survey.Serializer {
	return &surveyJSONSerializer{}
}

// Encode encodes a single survey into JSON format
func (s *surveyJSONSerializer) Encode(survey *survey.Survey) ([]byte, error) {
	return json.Marshal(survey)
}

// EncodeMultiple encodes multiple surveys into JSON format
func (s *surveyJSONSerializer) EncodeMultiple(surveys *survey.Surveys) ([]byte, error) {
	return json.Marshal(surveys)
}

// EncodeErrorResponse encodes an error response into JSON format
func (s *surveyJSONSerializer) EncodeErrorResponse(er survey.ErrorResponse) ([]byte, error) {
	return json.Marshal(er)
}

// Decode decodes a single survey from JSON format
func (s *surveyJSONSerializer) Decode(data []byte) (*survey.Survey, error) {
	survey := survey.Survey{}
	err := json.Unmarshal(data, &survey)
	return &survey, err
}

// DecodeMultiple decodes multiple surveys from JSON format
func (s *surveyJSONSerializer) DecodeMultiple(data []byte) (*survey.Surveys, error) {
	surveys := survey.Surveys{}
	err := json.Unmarshal(data, &surveys)
	return &surveys, err
}

// GetContentType returns the content type for JSON
func (s *surveyJSONSerializer) GetContentType() string {
	return "application/json"
}
