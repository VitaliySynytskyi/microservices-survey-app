package repository

import (
	"context"
	"fmt"

	"github.com/VitaliySynytskyi/microservices-survey-app/vote-service/config"
	"github.com/VitaliySynytskyi/microservices-survey-app/vote-service/vote"
	"github.com/jackc/pgx/v4"
)

// postgresResultsRepository implements the vote.ResultsRepository interface using PostgreSQL
type postgresResultsRepository struct {
	config     config.PostgresConfig
	connection *pgx.Conn
}

// NewPostgresResultsRepository creates a new Postgres vote results repository
// This function initializes a new PostgreSQL connection and returns a repository instance
func NewPostgresResultsRepository(cfg config.PostgresConfig) (vote.ResultsRepository, error) {
	p := &postgresResultsRepository{config: cfg}

	err := p.connect()
	if err != nil {
		return nil, err
	}

	return p, nil
}

// connect establishes a connection to the PostgreSQL database
func (p *postgresResultsRepository) connect() error {
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

// GetResults retrieves the results for a given survey ID from the PostgreSQL database
func (p *postgresResultsRepository) GetResults(surveyID string) (vote.Results, error) {
	// Extract query generation to a separate method (Extract Method refactoring)
	query := p.buildQuery()

	rows, err := p.connection.Query(context.Background(), query, surveyID)
	if err != nil {
		return vote.Results{}, err
	}
	defer rows.Close()

	results, err := p.processRows(rows, surveyID)
	if err != nil {
		return vote.Results{}, err
	}

	// Check if there are no votes, which could also mean the survey ID is invalid
	if results.UpdatedAt == 0 {
		return results, vote.ErrResultsNotFound
	}

	return results, rows.Err()
}

// buildQuery generates the SQL query for retrieving survey results
// Extract Method refactoring applied here
func (p *postgresResultsRepository) buildQuery() string {
	return fmt.Sprintf("SELECT question, votes, last_update FROM %s WHERE survey = $1", p.config.Tables.Results)
}

// processRows processes the result rows from the query
// Extract Method refactoring applied here
func (p *postgresResultsRepository) processRows(rows pgx.Rows, surveyID string) (vote.Results, error) {
	results := vote.Results{
		Survey: surveyID,
	}

	for rows.Next() {
		var question int
		var votes int
		var lastUpdate int64

		// Extract the data from the row
		err := rows.Scan(&question, &votes, &lastUpdate)
		if err != nil {
			return results, err
		}

		// Add the results for this question - convert Result to QuestionResults
		questionResults := vote.QuestionResults{
			Question:   question,
			TotalVotes: votes,
		}
		results.Results = append(results.Results, questionResults)

		// Update the last update time
		if results.UpdatedAt < lastUpdate {
			results.UpdatedAt = lastUpdate
		}
	}

	return results, rows.Err()
}
