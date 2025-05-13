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

// Encode encodes a survey with status into JSON format
func (s *surveyJSONSerializer) Encode(surveyWithStatus *survey.SurveyWithStatus) ([]byte, error) {
	return json.Marshal(surveyWithStatus)
}

// EncodeMultiple encodes multiple surveys with status into JSON format
func (s *surveyJSONSerializer) EncodeMultiple(surveysWithStatus []survey.SurveyWithStatus) ([]byte, error) {
	return json.Marshal(surveysWithStatus)
}

// EncodeSurvey encodes a single survey without status into JSON format
func (s *surveyJSONSerializer) EncodeSurvey(survey *survey.Survey) ([]byte, error) {
	return json.Marshal(survey)
}

// EncodeErrorResponse encodes an error response into JSON format
func (s *surveyJSONSerializer) EncodeErrorResponse(er survey.ErrorResponse) ([]byte, error) {
	return json.Marshal(er)
}

// EncodeSurveyResponse encodes a survey response into JSON format
func (s *surveyJSONSerializer) EncodeSurveyResponse(sr survey.SurveyResponse) ([]byte, error) {
	return json.Marshal(sr)
}

// EncodeSurveysResponse encodes a surveys response into JSON format
func (s *surveyJSONSerializer) EncodeSurveysResponse(sr survey.SurveysResponse) ([]byte, error) {
	return json.Marshal(sr)
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
