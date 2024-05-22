package handler

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/VitaliySynytskyi/microservices-survey-app/survey-service/middleware"
	"github.com/VitaliySynytskyi/microservices-survey-app/survey-service/survey"
	"github.com/go-chi/chi"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSurveyHTTPHandler_Get(t *testing.T) {
	log := zerolog.New(nil)
	mockService := new(MockSurveyService)
	handler := NewSurveyHTTPHandler(mockService, &log)

	t.Run("success", func(t *testing.T) {
		surveyID := "123"
		expectedSurvey := &survey.Survey{
			ID:        surveyID,
			Name:      "Test Survey",
			CreatedAt: time.Now().Unix(),
			Questions: []survey.Question{
				{ID: 1, Text: "Question 1"},
			},
		}

		mockService.On("LoadByID", surveyID).Return(expectedSurvey, nil)

		mockSerializer := new(MockSerializer)
		expectedJSON := []byte(`{"id":"123","name":"Test Survey","created_at":1234567890,"questions":[{"id":1,"text":"Question 1"}]}`)
		mockSerializer.On("Encode", expectedSurvey).Return(expectedJSON, nil)

		req := httptest.NewRequest("GET", "/surveys/123", nil)
		ctx := context.WithValue(req.Context(), middleware.SerializerKey, mockSerializer)
		req = req.WithContext(ctx)

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "123")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		w := httptest.NewRecorder()

		handler.Get(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
		assert.Equal(t, expectedJSON, w.Body.Bytes())

		mockService.AssertExpectations(t)
		mockSerializer.AssertExpectations(t)
	})
}

func TestSurveyHTTPHandler_Post(t *testing.T) {
	log := zerolog.New(nil)
	mockService := new(MockSurveyService)
	handler := NewSurveyHTTPHandler(mockService, &log)

	t.Run("success", func(t *testing.T) {
		newSurvey := &survey.Survey{
			ID:        "123",
			Name:      "New Survey",
			CreatedAt: time.Now().Unix(),
			Questions: []survey.Question{
				{ID: 1, Text: "Question 1"},
			},
		}

		requestBody := []byte(`{"id":"123","name":"New Survey","created_at":1234567890,"questions":[{"id":1,"text":"Question 1"}]}`)
		mockSerializer := new(MockSerializer)
		mockSerializer.On("Decode", mock.Anything).Return(newSurvey, nil)
		mockSerializer.On("Encode", newSurvey).Return(requestBody, nil)

		mockService.On("Insert", newSurvey).Return(nil)

		req := httptest.NewRequest("POST", "/surveys", bytes.NewReader(requestBody))
		ctx := context.WithValue(req.Context(), middleware.SerializerKey, mockSerializer)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		handler.Post(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
		assert.Equal(t, requestBody, w.Body.Bytes())

		mockService.AssertExpectations(t)
		mockSerializer.AssertExpectations(t)
	})

	t.Run("decode error", func(t *testing.T) {
		requestBody := []byte(`invalid json`)
		mockSerializer := new(MockSerializer)

		// Create a dummy survey object
		dummySurvey := &survey.Survey{}

		// Pass the dummy survey object to the Decode method
		mockSerializer.On("Decode", mock.Anything).Return(dummySurvey, errors.New("decode error"))
		mockSerializer.On("EncodeErrorResponse", survey.ErrorResponse{Error: http.StatusText(http.StatusInternalServerError)}).Return([]byte(`{"error":"Internal Server Error"}`), nil)

		req := httptest.NewRequest("POST", "/surveys", bytes.NewReader(requestBody))
		ctx := context.WithValue(req.Context(), middleware.SerializerKey, mockSerializer)
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		handler.Post(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
		assert.Equal(t, `{"error":"Internal Server Error"}`, w.Body.String())

		mockService.AssertExpectations(t)
		mockSerializer.AssertExpectations(t)
	})

}
