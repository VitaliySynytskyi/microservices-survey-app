package repository

import (
	"sort"
	"sync"

	"github.com/VitaliySynytskyi/microservices-survey-app/survey-service/survey"
)

// memorySurveyRepository implements the survey.Repository interface using an in-memory store
type memorySurveyRepository struct {
	surveys map[string]*survey.Survey
	mu      sync.RWMutex
}

// NewSurveyMemoryRepository creates a new in-memory survey repository
func NewSurveyMemoryRepository() (survey.Repository, error) {
	return &memorySurveyRepository{
		surveys: make(map[string]*survey.Survey),
	}, nil
}

// Insert adds a new survey to the in-memory store
func (r *memorySurveyRepository) Insert(s *survey.Survey) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.surveys[s.ID] = s
	return nil
}

// LoadByID retrieves a survey by its ID from the in-memory store
func (r *memorySurveyRepository) LoadByID(id string) (*survey.Survey, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if s, ok := r.surveys[id]; ok {
		return s, nil
	}
	return nil, survey.ErrNotFound
}

// Load retrieves all surveys from the in-memory store
func (r *memorySurveyRepository) Load() (*survey.Surveys, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Convert the map to a slice
	surveys := make(survey.Surveys, 0, len(r.surveys))
	for _, s := range r.surveys {
		surveys = append(surveys, s)
	}

	// Sort by creation time, newest first
	sort.Slice(surveys, func(i, j int) bool {
		return surveys[i].CreatedAt > surveys[j].CreatedAt
	})

	return &surveys, nil
}

// Update updates an existing survey in the in-memory store
func (r *memorySurveyRepository) Update(s *survey.Survey) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check if the survey exists
	if _, ok := r.surveys[s.ID]; !ok {
		return survey.ErrNotFound
	}

	// Update the survey
	r.surveys[s.ID] = s
	return nil
}

// Delete removes a survey by ID from the in-memory store
func (r *memorySurveyRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check if the survey exists
	if _, ok := r.surveys[id]; !ok {
		return survey.ErrNotFound
	}

	// Delete the survey
	delete(r.surveys, id)
	return nil
}
