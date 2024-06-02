package repository

import (
	"sort"
	"sync"
	"time"

	"github.com/VitaliySynytskyi/microservices-survey-app/survey-service/survey"
)

type surveyMemoryRepository struct {
	storage map[string]*survey.Survey
	mutex   *sync.RWMutex
}

// NewSurveyMemoryRepository creates a new survey repository that stores in memory
func NewSurveyMemoryRepository() (survey.Repository, error) {
	return &surveyMemoryRepository{
		storage: make(map[string]*survey.Survey),
		mutex:   &sync.RWMutex{},
	}, nil
}

// Insert adds a new survey to the in-memory storage
func (r *surveyMemoryRepository) Insert(s *survey.Survey) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.storage[s.ID] = s
	return nil
}

// LoadByID retrieves a survey by its ID from the in-memory storage
func (r *surveyMemoryRepository) LoadByID(id string) (*survey.Survey, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if survey, ok := r.storage[id]; ok {
		return survey, nil
	}
	return nil, survey.ErrNotFound
}

// Load retrieves all surveys from the in-memory storage
func (r *surveyMemoryRepository) Load() (*survey.Surveys, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	surveys := make(survey.Surveys, 0, len(r.storage))

	for _, item := range r.storage {
		surveys = append(surveys, item)
	}

	// Sort surveys by CreatedAt in descending order
	sort.Slice(surveys, func(i, j int) bool {
		return time.Unix(surveys[i].CreatedAt, 0).After(time.Unix(surveys[j].CreatedAt, 0))
	})

	return &surveys, nil
}
