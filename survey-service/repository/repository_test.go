package repository

import (
	"testing"
	"time"

	"github.com/VitaliySynytskyi/microservices-survey-app/survey-service/config"
	"github.com/VitaliySynytskyi/microservices-survey-app/survey-service/survey"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

// TestMemoryRepository tests the survey memory repository functions
func TestMemoryRepository(t *testing.T) {
	repo, err := NewSurveyMemoryRepository()
	assert.NoError(t, err, "Expected no error on creating memory repository")

	s := &survey.Survey{
		ID:        "1",
		Name:      "Test Survey",
		CreatedAt: time.Now().UTC().Unix(),
		Questions: []survey.Question{
			{ID: 1, Text: "What is your name?"},
			{ID: 2, Text: "How old are you?"},
		},
	}

	// Insert the survey
	err = repo.Insert(s)
	assert.NoError(t, err, "Expected no error on inserting survey")

	// Load the survey by ID
	loadedSurvey, err := repo.LoadByID("1")
	assert.NoError(t, err, "Expected no error on loading survey by ID")
	assert.Equal(t, s, loadedSurvey, "Expected loaded survey to match inserted survey")

	// Load all surveys
	surveys, err := repo.Load()
	assert.NoError(t, err, "Expected no error on loading all surveys")
	assert.Contains(t, *surveys, s, "Expected loaded surveys to contain the inserted survey")
}

// TestMongoRepository tests the survey MongoDB repository functions
func TestMongoRepository(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	cfg := config.MongoConfig{
		URL:     "mongodb://localhost:27017",
		DB:      "survey-db",
		Timeout: 10 * time.Second,
	}

	mt.Run("Insert and LoadByID", func(mt *mtest.T) {
		repo := &mongoSurveyRepository{
			client: mt.Client,
			config: cfg,
		}

		s := &survey.Survey{
			ID:        "1",
			Name:      "Test Survey",
			CreatedAt: time.Now().UTC().Unix(),
			Questions: []survey.Question{
				{ID: 1, Text: "What is your name?"},
				{ID: 2, Text: "How old are you?"},
			},
		}

		// Mock insert operation
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		err := repo.Insert(s)
		assert.NoError(t, err, "Expected no error on inserting survey")

		// Mock find operation
		first := mtest.CreateCursorResponse(1, "survey-db.surveys", mtest.FirstBatch, bson.D{
			{Key: "id", Value: "1"},
			{Key: "name", Value: "Test Survey"},
			{Key: "createdAt", Value: s.CreatedAt},
			{Key: "questions", Value: bson.A{
				bson.D{{Key: "id", Value: 1}, {Key: "text", Value: "What is your name?"}},
				bson.D{{Key: "id", Value: 2}, {Key: "text", Value: "How old are you?"}},
			}},
		})
		mt.AddMockResponses(first)

		// Load the survey by ID
		loadedSurvey, err := repo.LoadByID("1")
		assert.NoError(t, err, "Expected no error on loading survey by ID")
		assert.Equal(t, s, loadedSurvey, "Expected loaded survey to match inserted survey")
	})

	mt.Run("Load", func(mt *mtest.T) {
		repo := &mongoSurveyRepository{
			client: mt.Client,
			config: cfg,
		}

		s1 := &survey.Survey{
			ID:        "1",
			Name:      "Test Survey 1",
			CreatedAt: time.Now().UTC().Unix(),
			Questions: []survey.Question{
				{ID: 1, Text: "What is your name?"},
				{ID: 2, Text: "How old are you?"},
			},
		}

		s2 := &survey.Survey{
			ID:        "2",
			Name:      "Test Survey 2",
			CreatedAt: time.Now().UTC().Unix(),
			Questions: []survey.Question{
				{ID: 1, Text: "What is your favorite color?"},
				{ID: 2, Text: "What is your favorite food?"},
			},
		}

		// Mock find operation
		first := mtest.CreateCursorResponse(1, "survey-db.surveys", mtest.FirstBatch, bson.D{
			{Key: "id", Value: "1"},
			{Key: "name", Value: "Test Survey 1"},
			{Key: "createdAt", Value: s1.CreatedAt},
			{Key: "questions", Value: bson.A{
				bson.D{{Key: "id", Value: 1}, {Key: "text", Value: "What is your name?"}},
				bson.D{{Key: "id", Value: 2}, {Key: "text", Value: "How old are you?"}},
			}},
		})
		second := mtest.CreateCursorResponse(1, "survey-db.surveys", mtest.NextBatch, bson.D{
			{Key: "id", Value: "2"},
			{Key: "name", Value: "Test Survey 2"},
			{Key: "createdAt", Value: s2.CreatedAt},
			{Key: "questions", Value: bson.A{
				bson.D{{Key: "id", Value: 1}, {Key: "text", Value: "What is your favorite color?"}},
				bson.D{{Key: "id", Value: 2}, {Key: "text", Value: "What is your favorite food?"}},
			}},
		})
		killCursors := mtest.CreateCursorResponse(0, "survey-db.surveys", mtest.NextBatch)
		mt.AddMockResponses(first, second, killCursors)

		// Load all surveys
		surveys, err := repo.Load()
		assert.NoError(t, err, "Expected no error on loading all surveys")
		assert.Len(t, *surveys, 2, "Expected two surveys to be loaded")
		assert.Contains(t, *surveys, s1, "Expected loaded surveys to contain the first survey")
		assert.Contains(t, *surveys, s2, "Expected loaded surveys to contain the second survey")
	})
}
