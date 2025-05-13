package repository

import (
	"context"

	"github.com/VitaliySynytskyi/microservices-survey-app/survey-service/config"
	"github.com/VitaliySynytskyi/microservices-survey-app/survey-service/survey"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// mongoSurveyRepository implements the survey.Repository interface using MongoDB
type mongoSurveyRepository struct {
	client *mongo.Client
	config config.MongoConfig
}

// NewSurveyMongoRepository creates a new MongoDB survey repository
func NewSurveyMongoRepository(cfg config.MongoConfig) (survey.Repository, error) {
	repo := &mongoSurveyRepository{config: cfg}
	err := repo.connect()
	return repo, err
}

// connect establishes a connection to the MongoDB server
func (r *mongoSurveyRepository) connect() error {
	ctx, cancel := r.contextWithTimeout()
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(r.config.URL))
	if err != nil {
		return err
	}

	// Verify connection with a ping
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	r.client = client
	return nil
}

// contextWithTimeout creates a context with a timeout based on the repository configuration
func (r *mongoSurveyRepository) contextWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), r.config.Timeout)
}

// getCollection returns the MongoDB collection for surveys
func (r *mongoSurveyRepository) getCollection() *mongo.Collection {
	return r.client.Database(r.config.DB).Collection("surveys")
}

// Insert adds a new survey to the MongoDB collection
func (r *mongoSurveyRepository) Insert(s *survey.Survey) error {
	ctx, cancel := r.contextWithTimeout()
	defer cancel()

	_, err := r.getCollection().InsertOne(ctx, s)
	return err
}

// LoadByID retrieves a survey by its ID from the MongoDB collection
func (r *mongoSurveyRepository) LoadByID(id string) (*survey.Survey, error) {
	ctx, cancel := r.contextWithTimeout()
	defer cancel()

	s := &survey.Survey{}
	err := r.getCollection().FindOne(ctx, bson.M{"id": id}).Decode(s)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, survey.ErrNotFound
		}
		return nil, err
	}
	return s, nil
}

// Load retrieves all surveys from the MongoDB collection
func (r *mongoSurveyRepository) Load() (*survey.Surveys, error) {
	ctx, cancel := r.contextWithTimeout()
	defer cancel()

	surveys := make(survey.Surveys, 0)
	opts := options.Find().SetSort(bson.D{{Key: "createdAt", Value: -1}}).SetLimit(25)
	cursor, err := r.getCollection().Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &surveys)
	return &surveys, err
}

// Update updates an existing survey in the MongoDB collection
func (r *mongoSurveyRepository) Update(s *survey.Survey) error {
	ctx, cancel := r.contextWithTimeout()
	defer cancel()

	filter := bson.M{"id": s.ID}
	update := bson.M{"$set": s}

	result, err := r.getCollection().UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return survey.ErrNotFound
	}

	return nil
}

// Delete removes a survey by ID from the MongoDB collection
func (r *mongoSurveyRepository) Delete(id string) error {
	ctx, cancel := r.contextWithTimeout()
	defer cancel()

	filter := bson.M{"id": id}

	result, err := r.getCollection().DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return survey.ErrNotFound
	}

	return nil
}
