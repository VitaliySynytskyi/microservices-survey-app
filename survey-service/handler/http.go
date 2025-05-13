package handler

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/VitaliySynytskyi/microservices-survey-app/survey-service/middleware"
	"github.com/VitaliySynytskyi/microservices-survey-app/survey-service/survey"
	"github.com/go-chi/chi"
	"github.com/rs/zerolog"
)

// SurveyHTTPHandler handles HTTP requests for surveys
type SurveyHTTPHandler struct {
	service survey.Service
	log     *zerolog.Logger
}

// NewSurveyHTTPHandler creates a new survey HTTP handler
func NewSurveyHTTPHandler(service survey.Service, log *zerolog.Logger) *SurveyHTTPHandler {
	return &SurveyHTTPHandler{
		service,
		log,
	}
}

// Get handles get requests to get a single survey
func (h *SurveyHTTPHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	h.log.Info().Str("id", id).Msg("GET request received")

	// Load the requested survey
	surveyWithStatus, err := h.service.LoadByID(id)
	if err != nil {
		handleGetError(h, w, r, id, err)
		return
	}

	// Encode the survey to be returned
	serializer := h.GetSerializer(r)
	json, err := serializer.Encode(surveyWithStatus)
	if err != nil {
		h.log.Error().Str("id", id).Msg("Unable to encode survey")
		h.Error(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	h.Response(w, r, json, http.StatusOK)
}

// handleGetError handles errors during Get requests
// Code Smell: Long Method
// Refactoring: Extract Method to handle error responses for Get
func handleGetError(h *SurveyHTTPHandler, w http.ResponseWriter, r *http.Request, id string, err error) {
	if errors.Is(err, survey.ErrNotFound) {
		h.log.Debug().Str("id", id).Msg("Survey not found")
		h.Error(w, r, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	} else {
		h.log.Error().Err(err).Str("id", id).Msg("Unable to load survey")
		h.Error(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// Collection handles get requests to return a collection of surveys
func (h *SurveyHTTPHandler) Collection(w http.ResponseWriter, r *http.Request) {
	h.log.Info().Msg("COLLECTION request received")

	// Load all surveys
	surveysWithStatus, err := h.service.Load()
	if err != nil {
		h.log.Error().Err(err).Msg("Unable to load surveys")
		h.Error(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Encode the surveys to be returned
	serializer := h.GetSerializer(r)
	json, err := serializer.EncodeMultiple(surveysWithStatus)
	if err != nil {
		h.log.Error().Err(err).Msg("Unable to encode surveys")
		h.Error(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	h.Response(w, r, json, http.StatusOK)
}

// ActiveCollection handles get requests to return active surveys
func (h *SurveyHTTPHandler) ActiveCollection(w http.ResponseWriter, r *http.Request) {
	h.log.Info().Msg("ACTIVE COLLECTION request received")

	// Load active surveys
	activeSurveys, err := h.service.LoadActive()
	if err != nil {
		h.log.Error().Err(err).Msg("Unable to load active surveys")
		h.Error(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Encode the surveys to be returned
	serializer := h.GetSerializer(r)
	json, err := serializer.EncodeMultiple(activeSurveys)
	if err != nil {
		h.log.Error().Err(err).Msg("Unable to encode active surveys")
		h.Error(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	h.Response(w, r, json, http.StatusOK)
}

// Post handles post requests to create a survey
func (h *SurveyHTTPHandler) Post(w http.ResponseWriter, r *http.Request) {
	h.log.Info().Msg("POST request received")

	// Read the request body
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.log.Error().Err(err).Msg("Unable to read survey POST body")
		h.Error(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Decode the request
	serializer := h.GetSerializer(r)
	s, err := serializer.Decode(requestBody)
	if err != nil {
		h.log.Error().Err(err).Msg("Unable to decode survey POST body")
		h.Error(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Save the survey
	err = h.service.Insert(s)
	if err != nil {
		handlePostError(h, w, r, err)
		return
	}

	h.log.Info().Str("id", s.ID).Msg("Survey created")

	// Get survey with status
	surveyWithStatus, err := h.service.LoadByID(s.ID)
	if err != nil {
		h.log.Error().Str("id", s.ID).Err(err).Msg("Unable to load created survey")
		h.Error(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Encode the survey to be returned
	res, err := serializer.Encode(surveyWithStatus)
	if err != nil {
		h.log.Error().Str("id", s.ID).Err(err).Msg("Unable to encode survey")
		h.Error(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	h.Response(w, r, res, http.StatusCreated)
}

// Put handles put requests to update a survey
func (h *SurveyHTTPHandler) Put(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	h.log.Info().Str("id", id).Msg("PUT request received")

	// Read the request body
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.log.Error().Err(err).Msg("Unable to read survey PUT body")
		h.Error(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Decode the request
	serializer := h.GetSerializer(r)
	s, err := serializer.Decode(requestBody)
	if err != nil {
		h.log.Error().Err(err).Msg("Unable to decode survey PUT body")
		h.Error(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Update the survey
	err = h.service.Update(id, s)
	if err != nil {
		handleUpdateError(h, w, r, id, err)
		return
	}

	h.log.Info().Str("id", id).Msg("Survey updated")

	// Get updated survey with status
	updatedSurvey, err := h.service.LoadByID(id)
	if err != nil {
		h.log.Error().Str("id", id).Err(err).Msg("Unable to load updated survey")
		h.Error(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Encode the survey to be returned
	res, err := serializer.Encode(updatedSurvey)
	if err != nil {
		h.log.Error().Str("id", id).Err(err).Msg("Unable to encode updated survey")
		h.Error(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	h.Response(w, r, res, http.StatusOK)
}

// Delete handles delete requests to remove a survey
func (h *SurveyHTTPHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	h.log.Info().Str("id", id).Msg("DELETE request received")

	// Delete the survey
	err := h.service.DeleteSurvey(id)
	if err != nil {
		if errors.Is(err, survey.ErrNotFound) {
			h.log.Debug().Str("id", id).Msg("Survey not found for deletion")
			h.Error(w, r, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			h.log.Error().Err(err).Str("id", id).Msg("Unable to delete survey")
			h.Error(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	h.log.Info().Str("id", id).Msg("Survey deleted")
	h.Response(w, r, []byte(`{"message":"Survey deleted successfully"}`), http.StatusOK)
}

// Activate handles requests to activate a survey
func (h *SurveyHTTPHandler) Activate(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	h.log.Info().Str("id", id).Msg("ACTIVATE request received")

	// Activate the survey
	err := h.service.ActivateSurvey(id)
	if err != nil {
		if errors.Is(err, survey.ErrNotFound) {
			h.log.Debug().Str("id", id).Msg("Survey not found for activation")
			h.Error(w, r, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			h.log.Error().Err(err).Str("id", id).Msg("Unable to activate survey")
			h.Error(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	h.log.Info().Str("id", id).Msg("Survey activated")
	h.Response(w, r, []byte(`{"message":"Survey activated successfully"}`), http.StatusOK)
}

// Deactivate handles requests to deactivate a survey
func (h *SurveyHTTPHandler) Deactivate(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	h.log.Info().Str("id", id).Msg("DEACTIVATE request received")

	// Deactivate the survey
	err := h.service.DeactivateSurvey(id)
	if err != nil {
		if errors.Is(err, survey.ErrNotFound) {
			h.log.Debug().Str("id", id).Msg("Survey not found for deactivation")
			h.Error(w, r, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			h.log.Error().Err(err).Str("id", id).Msg("Unable to deactivate survey")
			h.Error(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	h.log.Info().Str("id", id).Msg("Survey deactivated")
	h.Response(w, r, []byte(`{"message":"Survey deactivated successfully"}`), http.StatusOK)
}

// SetExpiration handles requests to set an expiration date for a survey
func (h *SurveyHTTPHandler) SetExpiration(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	h.log.Info().Str("id", id).Msg("SET EXPIRATION request received")

	// Parse timestamp from query parameter
	expiresAtStr := r.URL.Query().Get("timestamp")
	if expiresAtStr == "" {
		h.log.Debug().Str("id", id).Msg("Missing timestamp parameter")
		h.Error(w, r, "Missing required timestamp parameter", http.StatusBadRequest)
		return
	}

	expiresAt, err := strconv.ParseInt(expiresAtStr, 10, 64)
	if err != nil {
		h.log.Debug().Str("id", id).Str("timestamp", expiresAtStr).Msg("Invalid timestamp format")
		h.Error(w, r, "Invalid timestamp format, must be Unix timestamp", http.StatusBadRequest)
		return
	}

	// Set the expiration date
	err = h.service.SetExpirationDate(id, expiresAt)
	if err != nil {
		if errors.Is(err, survey.ErrNotFound) {
			h.log.Debug().Str("id", id).Msg("Survey not found for setting expiration")
			h.Error(w, r, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else if errors.Is(err, survey.ErrInvalidRequest) {
			h.log.Debug().Str("id", id).Int64("timestamp", expiresAt).Msg("Invalid expiration timestamp")
			h.Error(w, r, err.Error(), http.StatusBadRequest)
		} else {
			h.log.Error().Err(err).Str("id", id).Msg("Unable to set survey expiration")
			h.Error(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	h.log.Info().Str("id", id).Int64("expiresAt", expiresAt).Msg("Survey expiration set")
	h.Response(w, r, []byte(`{"message":"Survey expiration set successfully"}`), http.StatusOK)
}

// handlePostError handles errors during Post requests
// Code Smell: Long Method
// Refactoring: Extract Method to handle error responses for Post
func handlePostError(h *SurveyHTTPHandler, w http.ResponseWriter, r *http.Request, err error) {
	if errors.Is(err, survey.ErrInvalidRequest) {
		h.log.Debug().Err(err).Msg("Invalid survey data in POST")
		h.Error(w, r, err.Error(), http.StatusUnprocessableEntity)
	} else if errors.Is(err, survey.ErrQuestionConditionalLogic) {
		h.log.Debug().Err(err).Msg("Invalid conditional logic in POST")
		h.Error(w, r, err.Error(), http.StatusUnprocessableEntity)
	} else {
		h.log.Error().Err(err).Msg("Unable to save survey")
		h.Error(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// handleUpdateError handles errors during Put requests
func handleUpdateError(h *SurveyHTTPHandler, w http.ResponseWriter, r *http.Request, id string, err error) {
	if errors.Is(err, survey.ErrNotFound) {
		h.log.Debug().Str("id", id).Msg("Survey not found for update")
		h.Error(w, r, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	} else if errors.Is(err, survey.ErrInvalidRequest) {
		h.log.Debug().Err(err).Str("id", id).Msg("Invalid survey data in PUT")
		h.Error(w, r, err.Error(), http.StatusUnprocessableEntity)
	} else if errors.Is(err, survey.ErrQuestionConditionalLogic) {
		h.log.Debug().Err(err).Str("id", id).Msg("Invalid conditional logic in PUT")
		h.Error(w, r, err.Error(), http.StatusUnprocessableEntity)
	} else {
		h.log.Error().Err(err).Str("id", id).Msg("Unable to update survey")
		h.Error(w, r, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// GetSerializer gets a serializer from the request, which is added via middleware
func (h *SurveyHTTPHandler) GetSerializer(r *http.Request) survey.Serializer {
	return r.Context().Value(middleware.SerializerKey).(survey.Serializer)
}

// Response sends an HTTP response
func (h *SurveyHTTPHandler) Response(w http.ResponseWriter, r *http.Request, output []byte, code int) {
	w.Header().Set("Content-Type", h.GetSerializer(r).GetContentType())
	w.WriteHeader(code)
	w.Write(output)
}

// Error sends an HTTP error response
func (h *SurveyHTTPHandler) Error(w http.ResponseWriter, r *http.Request, message string, code int) {
	serializer := h.GetSerializer(r)
	output, err := serializer.EncodeErrorResponse(survey.ErrorResponse{Error: message})

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	h.Response(w, r, output, code)
}
