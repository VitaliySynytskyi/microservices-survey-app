package handler

import (
	"context"
	"errors"
	"testing"
	"time"

	protos "github.com/VitaliySynytskyi/microservices-survey-app/survey-service/protos/survey"
	"github.com/VitaliySynytskyi/microservices-survey-app/survey-service/survey"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestGetSurvey(t *testing.T) {
	log := zerolog.New(nil)

	t.Run("success", func(t *testing.T) {
		mockService := new(MockSurveyService)
		handler := NewSurveyGrpcHandler(mockService, &log)

		surveyID := "123"
		expectedSurvey := &survey.Survey{
			ID:        surveyID,
			Name:      "Test Survey",
			CreatedAt: time.Now().Unix(), // CreatedAt має бути int64
			Questions: []survey.Question{
				{ID: 1, Text: "Question 1"},
			},
		}

		mockService.On("LoadByID", surveyID).Return(expectedSurvey, nil)

		req := &protos.SurveyRequest{Id: surveyID}
		res, err := handler.GetSurvey(context.Background(), req)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, surveyID, res.Id)
		assert.Equal(t, "Test Survey", res.Name)
		assert.Equal(t, 1, len(res.Questions))
		assert.Equal(t, int32(1), res.Questions[0].Id)
		assert.Equal(t, "Question 1", res.Questions[0].Text)

		mockService.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		mockService := new(MockSurveyService)
		handler := NewSurveyGrpcHandler(mockService, &log)

		surveyID := "123"
		mockService.On("LoadByID", surveyID).Return(nil, survey.ErrNotFound)

		req := &protos.SurveyRequest{Id: surveyID}
		res, err := handler.GetSurvey(context.Background(), req)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.True(t, errors.Is(err, survey.ErrNotFound))

		mockService.AssertExpectations(t)
	})

	t.Run("internal error", func(t *testing.T) {
		mockService := new(MockSurveyService)
		handler := NewSurveyGrpcHandler(mockService, &log)

		surveyID := "123"
		mockService.On("LoadByID", surveyID).Return(nil, errors.New("internal error"))

		req := &protos.SurveyRequest{Id: surveyID}
		res, err := handler.GetSurvey(context.Background(), req)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "internal error", err.Error())

		mockService.AssertExpectations(t)
	})
}
