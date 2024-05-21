package storage

import "github.com/VitaliySynytskyi/microservices-survey-app/vote-service/vote"

// VoteStorage contains functions to store votes and the results of votes
// This interface defines the methods required for storing votes and updating vote results
type VoteStorage interface {
	// Insert inserts a vote into storage
	// This method saves a new vote to the storage
	Insert(v *vote.Vote) error

	// UpdateResults updates the total vote count results for the survey and
	// question of a given vote
	// This method updates the vote count for a specific survey and question
	UpdateResults(v *vote.Vote) error
}
