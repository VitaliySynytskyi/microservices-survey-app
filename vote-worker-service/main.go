package main

import (
	"os"

	"github.com/VitaliySynytskyi/microservices-survey-app/vote-service/serializer"
	"github.com/VitaliySynytskyi/microservices-survey-app/vote-service/vote"
	"github.com/VitaliySynytskyi/microservices-survey-app/vote-worker-service/config"
	"github.com/VitaliySynytskyi/microservices-survey-app/vote-worker-service/logger"
	"github.com/VitaliySynytskyi/microservices-survey-app/vote-worker-service/queue"
	"github.com/VitaliySynytskyi/microservices-survey-app/vote-worker-service/storage"
)

// TODO: Need to handle if DB or queue connection dies?

func main() {
	// Get a logger
	log := logger.NewConsoleLogger()

	// Load application configuration
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot load application configuration")
		os.Exit(1)
	}

	// Get a vote serializer
	sz := serializer.NewVoteJSONSerializer()

	// Load the storage
	stg, err := storage.NewPostgresVoteStorage(cfg.Postgres)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot connect to vote storage")
		os.Exit(1)
	}

	// Load the queue
	mq := queue.NewRabbitVoteQueue(cfg.Rabbit, sz, &log)

	// Create a channel to receive votes from the queue
	vc := make(chan *vote.Vote)

	// Consume the queue
	go func() {
		mq.Consume(vc)
	}()

	// Receive from the queue
	go func() {
		for v := range vc {
			// Store the vote
			err := stg.Insert(v)
			if err != nil {
				log.Error().Err(err).Str("id", v.ID).Msg("Unable to store vote")
				continue
			}
			log.Info().Str("id", v.ID).Msg("Vote stored")

			// Update the results
			err = stg.UpdateResults(v)
			if err != nil {
				log.Error().Err(err).Str("id", v.ID).Msg("Unable to update vote results")
				continue
			}
			log.Info().Str("id", v.ID).Msg("Vote added to results")
		}
	}()

	// Wait forever
	forever := make(chan bool)
	<-forever
}
