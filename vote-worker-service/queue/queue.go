package queue

import "github.com/VitaliySynytskyi/microservices-survey-app/vote-service/vote"

// VoteQueue contains functions to receive votes from a queue
// This interface defines the methods required for consuming votes from a queue
type VoteQueue interface {
	// Consume consumes votes from a queue and passes them through a channel to be processed
	// This method reads votes from the queue and sends them to the provided channel
	Consume(vc chan<- *vote.Vote)
}
