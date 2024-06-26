package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/VitaliySynytskyi/microservices-survey-app/vote-service/vote"
	"github.com/VitaliySynytskyi/microservices-survey-app/vote-worker-service/config"
	"github.com/jackc/pgx/v4"
)

// postgresVoteStorage implements the VoteStorage interface for storing votes in PostgreSQL
type postgresVoteStorage struct {
	config     config.PostgresConfig
	connection *pgx.Conn
}

// NewPostgresVoteStorage creates a new Postgres vote storage
// This function initializes and returns a new instance of postgresVoteStorage
func NewPostgresVoteStorage(cfg config.PostgresConfig) (VoteStorage, error) {
	p := &postgresVoteStorage{config: cfg}

	err := p.connect()
	if err != nil {
		return nil, err
	}

	return p, nil
}

// connect establishes a connection to the PostgreSQL database
// This method connects to the database using the provided configuration
func (p *postgresVoteStorage) connect() error {
	addr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		p.config.User,
		p.config.Password,
		p.config.Hostname,
		p.config.Port,
		p.config.Database,
	)
	conn, err := pgx.Connect(context.Background(), addr)
	if err != nil {
		return err
	}

	p.connection = conn
	return nil
}

// Insert inserts a new vote into the votes table
// This method saves the vote details into the PostgreSQL database
func (p *postgresVoteStorage) Insert(v *vote.Vote) error {
	q := fmt.Sprintf("INSERT INTO %s(id, survey, question, created) VALUES($1, $2, $3, $4)", p.config.Tables.Votes)
	_, err := p.connection.Exec(context.Background(), q, v.ID, v.Survey, v.Question, v.Timestamp)
	return err
}

// UpdateResults updates the vote results in the results table
// This method increments the vote count for a specific survey and question, or initializes it if not present
func (p *postgresVoteStorage) UpdateResults(v *vote.Vote) error {
	// Check if results already exist for the vote's survey and question
	var r int
	q := fmt.Sprintf("SELECT votes FROM %s WHERE survey = $1 AND question = $2", p.config.Tables.Results)
	err := p.connection.QueryRow(context.Background(), q, v.Survey, v.Question).Scan(&r)

	switch err {
	case nil:
		// Increment the results
		return p.incrementResults(v)
	case pgx.ErrNoRows:
		// Initialize the results
		return p.initializeResults(v)
	default:
		return err
	}
}

// incrementResults increments the vote count for a specific survey and question
func (p *postgresVoteStorage) incrementResults(v *vote.Vote) error {
	q := fmt.Sprintf("UPDATE %s SET votes = votes + 1, last_update = $1 WHERE survey = $2 AND question = $3", p.config.Tables.Results)
	_, err := p.connection.Exec(context.Background(), q, time.Now().UTC().Unix(), v.Survey, v.Question)
	return err
}

// initializeResults initializes the vote count for a specific survey and question
func (p *postgresVoteStorage) initializeResults(v *vote.Vote) error {
	q := fmt.Sprintf("INSERT INTO %s(survey, question, votes, last_update) VALUES($1, $2, $3, $4)", p.config.Tables.Results)
	_, err := p.connection.Exec(context.Background(), q, v.Survey, v.Question, 1, time.Now().UTC().Unix())
	return err
}
