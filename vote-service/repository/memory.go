package repository

import (
	"sync"

	"github.com/VitaliySynytskyi/microservices-survey-app/vote-service/vote"
)

type voteMemoryRepository struct {
	storage map[string]*vote.Vote
	mutex   *sync.RWMutex
}

// NewMemoryVoteWriterRepository creates a new vote writer repository that stores in memory
// This function initializes and returns a new in-memory vote repository
func NewMemoryVoteWriterRepository() (vote.WriterRepository, error) {
	return &voteMemoryRepository{
		storage: make(map[string]*vote.Vote),
		mutex:   &sync.RWMutex{},
	}, nil
}

// Insert adds a new vote to the in-memory storage
// This method locks the storage, inserts the vote, and then unlocks the storage
func (r *voteMemoryRepository) Insert(v *vote.Vote) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.storage[v.ID] = v
	return nil
}
