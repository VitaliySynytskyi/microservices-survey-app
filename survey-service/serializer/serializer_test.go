package serializer

import (
	"testing"

	"github.com/VitaliySynytskyi/microservices-survey-app/survey-service/survey"
	"github.com/stretchr/testify/assert"
)

func TestNewSurveyJSONSerializer(t *testing.T) {
	serializer := NewSurveyJSONSerializer()
	assert.NotNil(t, serializer, "Expected new serializer to be non-nil")
	assert.IsType(t, &surveyJSONSerializer{}, serializer, "Expected serializer to be of type surveyJSONSerializer")
}

func TestSurveyJSONSerializer_Encode(t *testing.T) {
	serializer := NewSurveyJSONSerializer()
	s := &survey.Survey{
		ID:   "1",
		Name: "Test Survey",
		Questions: []survey.Question{
			{ID: 1, Text: "What is your name?"},
			{ID: 2, Text: "How old are you?"},
		},
	}

	data, err := serializer.Encode(s)
	assert.NoError(t, err, "Expected no error on encoding survey")
	assert.JSONEq(t, `{"id":"1","name":"Test Survey","questions":[{"id":1,"text":"What is your name?"},{"id":2,"text":"How old are you?"}],"createdAt":0}`, string(data), "Expected JSON to match")
}

func TestSurveyJSONSerializer_EncodeMultiple(t *testing.T) {
	serializer := NewSurveyJSONSerializer()
	surveys := &survey.Surveys{
		&survey.Survey{
			ID:   "1",
			Name: "Test Survey 1",
			Questions: []survey.Question{
				{ID: 1, Text: "What is your name?"},
				{ID: 2, Text: "How old are you?"},
			},
		},
		&survey.Survey{
			ID:   "2",
			Name: "Test Survey 2",
			Questions: []survey.Question{
				{ID: 1, Text: "What is your favorite color?"},
				{ID: 2, Text: "What is your favorite food?"},
			},
		},
	}

	data, err := serializer.EncodeMultiple(surveys)
	assert.NoError(t, err, "Expected no error on encoding multiple surveys")
	expectedJSON := `[{"id":"1","name":"Test Survey 1","questions":[{"id":1,"text":"What is your name?"},{"id":2,"text":"How old are you?"}],"createdAt":0},{"id":"2","name":"Test Survey 2","questions":[{"id":1,"text":"What is your favorite color?"},{"id":2,"text":"What is your favorite food?"}],"createdAt":0}]`
	assert.JSONEq(t, expectedJSON, string(data), "Expected JSON to match")
}

func TestSurveyJSONSerializer_EncodeErrorResponse(t *testing.T) {
	serializer := NewSurveyJSONSerializer()
	errResp := survey.ErrorResponse{
		Error: "Some error occurred",
	}

	data, err := serializer.EncodeErrorResponse(errResp)
	assert.NoError(t, err, "Expected no error on encoding error response")
	assert.JSONEq(t, `{"error":"Some error occurred"}`, string(data), "Expected JSON to match")
}

func TestSurveyJSONSerializer_Decode(t *testing.T) {
	serializer := NewSurveyJSONSerializer()
	data := `{"id":"1","name":"Test Survey","questions":[{"id":1,"text":"What is your name?"},{"id":2,"text":"How old are you?"}],"createdAt":0}`
	s, err := serializer.Decode([]byte(data))
	assert.NoError(t, err, "Expected no error on decoding survey")
	expectedSurvey := &survey.Survey{
		ID:   "1",
		Name: "Test Survey",
		Questions: []survey.Question{
			{ID: 1, Text: "What is your name?"},
			{ID: 2, Text: "How old are you?"},
		},
	}
	assert.Equal(t, expectedSurvey, s, "Expected decoded survey to match")
}

func TestSurveyJSONSerializer_DecodeMultiple(t *testing.T) {
	serializer := NewSurveyJSONSerializer()
	data := `[{"id":"1","name":"Test Survey 1","questions":[{"id":1,"text":"What is your name?"},{"id":2,"text":"How old are you?"}],"createdAt":0},{"id":"2","name":"Test Survey 2","questions":[{"id":1,"text":"What is your favorite color?"},{"id":2,"text":"What is your favorite food?"}],"createdAt":0}]`
	surveys, err := serializer.DecodeMultiple([]byte(data))
	assert.NoError(t, err, "Expected no error on decoding multiple surveys")
	expectedSurveys := &survey.Surveys{
		&survey.Survey{
			ID:   "1",
			Name: "Test Survey 1",
			Questions: []survey.Question{
				{ID: 1, Text: "What is your name?"},
				{ID: 2, Text: "How old are you?"},
			},
		},
		&survey.Survey{
			ID:   "2",
			Name: "Test Survey 2",
			Questions: []survey.Question{
				{ID: 1, Text: "What is your favorite color?"},
				{ID: 2, Text: "What is your favorite food?"},
			},
		},
	}
	assert.Equal(t, expectedSurveys, surveys, "Expected decoded surveys to match")
}

func TestSurveyJSONSerializer_GetContentType(t *testing.T) {
	serializer := NewSurveyJSONSerializer()
	contentType := serializer.GetContentType()
	assert.Equal(t, "application/json", contentType, "Expected content type to be application/json")
}
